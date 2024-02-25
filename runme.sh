#!/bin/bash

#Simple ejemplo de parametros necesarios:
PATH_FILEPASSW=rockyou.txt
SERVER=192.168.111.104
USER=admin
PORT=2222
THREADS=10

go run main.go -h ${SERVER} -p ${PORT} -u ${USER} -t ${THREADS}
