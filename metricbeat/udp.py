import socket
import json

UDP_IP = "127.0.0.1"
UDP_PORT = 8080
MESSAGE = "Hello, World!"
EVENT = {}
EVENT["text"] = "Hello World"

print "UDP target IP:", UDP_IP
print "UDP target port:", UDP_PORT
print "message:", MESSAGE

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM) # UDP

for n in range(0, 100):
    EVENT["metric"] = n
    sock.sendto(json.dumps(EVENT), (UDP_IP, UDP_PORT))
