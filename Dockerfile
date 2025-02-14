FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

ENV GO111MODULE=on
ENV GOSUMDB=off
ENV GOPROXY=https://goproxy.cn

RUN go mod tidy && go build -o feishuBot .
RUN chmod +x feishuBot

FROM busybox:1.37.0-glibc

WORKDIR /root/

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /root/time/zoneinfo.zip
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /app/feishuBot .

COPY entrypoint.sh /root/entrypoint.sh

ENTRYPOINT ["sh", "/root/entrypoint.sh"]
