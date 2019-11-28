package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

// CGO_ENABLED: 0
// GOOS: darwin、freebsd、linux、windows
// GOARCH: 386、amd64、arm
//
//-go:generate bash -c "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/kgm_cracker kgm_cracker.go"
//go:generate bash -c "CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/kgm_cracker.exe kgm_cracker.go"
//go:generate bash -c "CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/kgm_cracker_x86.exe kgm_cracker.go"

type Handler struct {
}

func (h Handler) ServeHTTP(rep http.ResponseWriter, req *http.Request) {
	fmt.Printf("收到验证请求，已返回成功激活数据\n")
	rep.Write([]byte("{\"code\":0,\"message\":\"破解成功\"}"))
}

func main() {
	var handler Handler
	fmt.Printf("Please add\n\t127.0.0.1 yinyuezhushou.com\nto your hosts file (C:\\Windows\\System32\\drivers\\etc\\hosts)\n\n\n")
	
	switch runtime.GOOS {
	case "windows":
		exec.Command("explorer.exe", "C:\\Windows\\System32\\drivers\\etc\\").Run()
	default:
		fmt.Printf("Maybe you should run on Windows\n")
	}

	err := http.ListenAndServe("127.0.0.1:8008", handler)
	if err != nil {
		panic(err)
	}
}
