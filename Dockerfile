FROM golang:1.18-alpine as go-builder

RUN apk --no-cache --no-progress add ca-certificates git tzdata  \
    && update-ca-certificates \
    && rm -rf /var/cache/apk/*

WORKDIR /go/work

COPY go.mod ./  go.sum ./ 
COPY cloudapp.go ./ 
COPY handler ./handler/ 
COPY model ./model/
COPY vender ./vender/ 
COPY web ./web/

# Download go modules
ENV GOPROXY="https://proxy.golang.com.cn,direct"
ENV GO111MODULE=on

RUN go mod download
#RUN GO111MODULE=on GOPROXY=https://goproxy.cn,direct go mod download
#RUN GO111MODULE=on GOPROXY=https://proxy.golang.com.cn,direct go mod download

#RUN make build
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build cloud
  #  -ldflags "-X ${REPO_PATH}/pkg/version.Version=${VERSION} -X ${REPO_PATH}/pkg/version.GitSHA=${GIT_SHA}" \
  #  $BUILD_PATH

# -------------------------------------------

FROM alpine:3.9 AS final

WORKDIR /work
RUN set -eux; \
	echo 'Asia/Shanghai' >/etc/timezone;

COPY --from=go-builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=go-builder /go/work/cloud /work/

ENTRYPOINT ["/work/cloud"]
EXPOSE 80
