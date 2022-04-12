set GOOS=linux
set CGO_ENABLE=0
set GOARCH=amd64
go build -o apiproxy myproxy.go


CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o apiproxy myproxy.go
