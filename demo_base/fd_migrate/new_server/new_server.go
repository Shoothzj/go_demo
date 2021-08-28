package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"
)

func main() {
	// 新版本sidecar 接收ListenFD，并且开始对外服务
	// 监听UDS
	addr, _ := net.ResolveUnixAddr("unix", "/dev/shm/migrate.sock")
	unixLn, _ := net.ListenUnix("unix", addr)
	conn, _ := unixLn.AcceptUnix()
	buf := make([]byte, 32)
	oob := make([]byte, 32)
	time.Sleep(5 * time.Second)
	// 接收 ListenFD
	_, oobn, _, _, _ := conn.ReadMsgUnix(buf, oob)
	scms, _ := syscall.ParseSocketControlMessage(oob[:oobn])
	if len(scms) > 0 {
		// 解析FD，并转化为 *net.TCPListener
		fds, _ := syscall.ParseUnixRights(&scms[0])
		f := os.NewFile(uintptr(fds[0]), "")
		ln, _ := net.FileListener(f)
		tcpLn, _ := ln.(*net.TCPListener)
		http.Serve(tcpLn, &MyHttp{})
	}
}

type MyHttp struct {
}

func (*MyHttp) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world")
}
