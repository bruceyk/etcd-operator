FROM golang:1.24.5-alpine

RUN apk add --no-cache ca-certificates

ADD _output/bin/etcd-backup-operator /usr/local/bin/etcd-backup-operator
ADD _output/bin/etcd-restore-operator /usr/local/bin/etcd-restore-operator
ADD _output/bin/etcd-operator /usr/local/bin/etcd-operator

RUN adduser -D etcd-operator
USER etcd-operator
