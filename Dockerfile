FROM golang:1.16 as build

WORKDIR /go/src/github.com/deployKubernetesInCHINA/dkc-command
COPY . .

RUN export GOPROXY=https://goproxy.cn && go mod download && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dkc-linux-amd64 .

FROM alpine

RUN apk --no-cache add curl

LABEL user=wenchangcui

RUN apk  add --no-cache tzdata
COPY --from=build /go/src/github.com/deployKubernetesInCHINA/dkc-command/dkc-linux-amd64 /dkc/
ENV TZ=Asia/Shanghai
EXPOSE 5555
WORKDIR /dkc
COPY README.md .
COPY docs/ docs
COPY inventory/ inventory/
COPY static/ static/
COPY views/ views/
COPY kubespray/ kubespray/
COPY union_rsa .
ENTRYPOINT ["/dkc/dkc-linux-amd64", "web"]
