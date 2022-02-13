

buildServer:
	go build -o server.exe cmd/worldtest/main.go



default: buildServer