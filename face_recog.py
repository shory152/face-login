#!/usr/bin/env python
# -*- coding:utf-8 -*-
import socket
import sys
import matplotlib.pyplot as plt
from PIL import Image
import time
import demjson

ip_port = ('127.0.0.1',9999)
sk = socket.socket()
sk.connect(ip_port)

def test_search(photo):
    str='{"cmd":"search","pic":"' + photo + '"}'
    str_len =len(str)
    send_str="%04d"%str_len + str
    return send_str

def test_add_face(id, photo):
    str = '{"cmd":"add_index","id":' + id + ',"pic":"' + photo + '"}'
    str_len = len(str)
    send_str = "%04d" % str_len + str  
    return send_str

def test_detect_face(photo):
    str = '{"cmd":"face_detect","pic":"' + photo + '"}'
    str_len = len(str)
    send_str = "%04d" % str_len + str
    return send_str

if len(sys.argv) < 2:
	print("missing arguments!, please give id and photo")
	exit()

print("args: ", sys.argv)
imgfn = sys.argv[1]
btim = time.time()

send_str=test_search(imgfn)
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

etim = time.time()
elapse = etim - btim
print("elapse %d seconds" % elapse)

rs = demjson.decode(server_reply)
if rs is None:
    print("no result\n")
    exit()
if rs["data"][1][0] is None:
    print("no distance\n")
    exit()
if rs['data'][1][0] >= 0.6:
    print("not recognized\n")

id = rs['data'][0][0]

fn=""
f = open('idface.txt')
for i in f:
    idfn = str.split(i,':')
    if int(idfn[0]) == id:
        fn = idfn[1]
        break
f.close()

if fn == "":
    print("fn not found")
    exit()
fn = fn.rstrip('\n')

img1 = Image.open(imgfn)
img2 = Image.open(fn)

plt.figure()
sub1 = plt.subplot(1,2,1)
sub1.imshow(img1)
sub2 = plt.subplot(1,2,2)
sub2.imshow(img2)
plt.show()




