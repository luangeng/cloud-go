  
.PHONY: default build image check publish-images

default: build

build:
		CGO_ENABLED=0 GO111MODULE=off go build -a --trimpath --installsuffix cgo --ldflags="-s" -o cloud

test:
		go test -v -cover ./...
