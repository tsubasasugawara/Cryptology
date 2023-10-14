FROM golang:alpine

RUN apk update && apk add vim curl bash git openssh-keygen

WORKDIR /go/src
COPY . .

CMD ["bash"]
