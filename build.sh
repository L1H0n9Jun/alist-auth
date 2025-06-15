GOOS=linux GOARCH=amd64 GOAMD64=v3 \
go build -trimpath -ldflags '-w -s -buildid=' \
-o ali-openapi
