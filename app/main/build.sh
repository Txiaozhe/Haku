#! /bin/sh

GOOS=linux GOARCH=amd64 go build

scp main root@106.15.227.154:~/workspace/haku
