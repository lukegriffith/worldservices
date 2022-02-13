

buildServer:
	go build -o server.exe cmd/worldtest/main.go

test:
	go test


default: buildServer
