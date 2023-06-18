FROM golang:1.18

# ENV CGO_ENABLED 0
# ENV GOOS linux
# ENV GOPROXY https://goproxy.cn,direct

WORKDIR /app

ADD httpserver/go.mod ./
# ADD httpserver/go.sum ./
RUN go mod download
COPY httpserver/ ./

RUN go build -ldflags="-s -w" -o ./server main.go

CMD /server

CMD ["./server"]