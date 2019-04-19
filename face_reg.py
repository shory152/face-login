#!/usr/bin/env python
# -*- coding:utf-8 -*-
import socket
import sys



def test_add_face(id, photo):
    str = '{"cmd":"add_index","id":' + id + ',"pic":"' + photo + '"}'
    str_len = len(str)
    send_str = "%04d" % str_len + str  
    return send_str

if len(sys.argv) <= 2:
	print("missing arguments!, please give id and photo")
	exit()

print("args: ", sys.argv)

ip_port = ('127.0.0.1',9999)
sk = socket.socket()
sk.connect(ip_port)

send_str=test_add_face(sys.argv[1], sys.argv[2])
print("send:", send_str)
send_str = bytes(send_str, encoding="utf8")
sk.sendall(send_str)

server_reply = sk.recv(4)
server_reply = str(server_reply, encoding="utf8")
print (server_reply+"\n")

server_reply = int(server_reply)
server_reply=sk.recv(server_reply);
server_reply = str(server_reply, encoding="utf8")
print (server_reply+"\n")

sk.close()

