FROM golang:1-alpine as builder

RUN apk --no-cache --no-progress add git ca-certificates tzdata make \
    && update-ca-certificates \
    && rm -rf /var/cache/apk/*

WORKDIR /go/tool

# Download go modules
COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on GOPROXY=https://proxy.golang.org go mod download

COPY upx .
RUN chmod a+x upx
COPY app.go .
COPY Makefile . 

RUN make build
RUN ./upx -9 -o ./tool ./tool1

FROM debian:buster-slim

RUN set -eux; \
	apt-get update; \
	apt-get install -y --no-install-recommends less vim curl wget jq tree procps net-tools iputils-ping --fix-missing; \
	rm -rf /var/lib/apt/lists/*; \
	rm -rf /usr/share/zoneinfo/*; \
	echo 'Asia/Shanghai' >/etc/timezone;

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=builder /go/tool/tool .

ENTRYPOINT ["/tool"]
EXPOSE 80
