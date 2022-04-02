FROM golang:1-alpine as builder

RUN apk --no-cache --no-progress add git ca-certificates tzdata make \
    && update-ca-certificates \
    && rm -rf /var/cache/apk/*

WORKDIR /go/work

# Download go modules
#ENV GOPROXY="https://proxy.golang.com.cn,direct"
COPY go.mod .
COPY go.sum .
#RUN GO111MODULE=on GOPROXY=https://proxy.golang.org go mod download
RUN GO111MODULE=on GOPROXY=https://proxy.golang.com.cn,direct go mod download

#COPY upx .
#RUN chmod a+x upx
COPY * ./
COPY Makefile . 

RUN make build
#RUN ./upx -9 -o ./cloud ./cloud1

FROM debian:buster-slim

RUN set -eux; \
	apt-get update; \
	apt-get install -y --no-install-recommends less vim curl wget jq tree procps net-tools iputils-ping --fix-missing; \
	rm -rf /var/lib/apt/lists/*; \
	rm -rf /usr/share/zoneinfo/*; \
	echo 'Asia/Shanghai' >/etc/timezone;

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=builder /go/work/cloud .

ENTRYPOINT ["/cloud"]
EXPOSE 80
