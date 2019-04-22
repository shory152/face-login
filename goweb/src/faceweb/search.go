package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func searchFace(w http.ResponseWriter, req *http.Request) {
	b64img := req.PostFormValue("img")
	if b64img == "" {
		replyErr(w, "not found image")
		return
	}

	img, err := base64.StdEncoding.DecodeString(b64img)
	if err != nil {
		replyErr(w, "invalid image")
		return
	}

	fn := fmt.Sprintf("recvimg/%012d.png", time.Now().UnixNano())
	ioutil.WriteFile(fn, img, 0677)

	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		replyErr(w, "can not connect to face search server")
		return
	}
	defer conn.Close()

	/*
			str='{"cmd":"search","pic":"' + photo + '"}'
		    str_len =len(str)
		    send_str="%04d"%str_len + str
	*/
	cmdstr := `{"cmd":"search","pic":"%s"}`
	cmdstr = fmt.Sprintf(cmdstr, fn)
	cmdlen := fmt.Sprintf("%04d", len(cmdstr))
	cmdstr = cmdlen + cmdstr

	io.WriteString(conn, cmdstr)

	buf := make([]byte, 1048576)
	replen, err := conn.Read(buf)
	if err != nil {
		replyErr(w, err.Error())
		return
	}
	if replen <= 0 {
		replyErr(w, "recognize fail, no result returned")
		return
	}

	repbody, err := decodeFace(string(buf[4 : replen-4]))
	if err != nil {
		replyErr(w, err.Error())
		return
	}

	if repbody.code != 0 {
		replyErr(w, "recognize error")
		return
	}

	if repbody.data[1].(disary)[0] < 0.6 {
		id := repbody.data[0].(idary)[0]
		replyOK(w, getReplyFn(id))
		return
	}

}

func getReplyFn(id int) string {
	if id < 1000 {
		return fmt.Sprintf("../staffImage/B00%3d.jpg", id)
	} else {
		return fmt.Sprintf("../newface/%d.jpg", id)
	}
}
func replyErr(w http.ResponseWriter, msg string) {
	io.WriteString(w, `{"code":1,"error":"`+msg+`"}`)
}
func replyOK(w http.ResponseWriter, fn string) {
	io.WriteString(w, `{"code":0,"data":{"file_path":"`+fn+`"}}`)
}

/*
	send: 0037{"cmd":"search","pic":"photo/s6.jpg"}
0150

{'code': 0, 'data': [[1068, 167, 296, 115, 375],
 [0.5269452929496765, 0.6523714661598206, 0.6586498022079468, 0.6722832322120667, 0.677651047706604]]}
*/
type idary [5]int
type disary [5]float64
type iddis [2]interface{}
type facerep struct {
	code int
	data iddis
}

func decodeFace(jsonstr string) (*facerep, error) {
	jdec := json.NewDecoder(bytes.NewReader([]byte(jsonstr)))
	var rep facerep
	rep.data = iddis{&idary{}, &disary{}}
	err := jdec.Decode(&rep)
	return &rep, err
}
