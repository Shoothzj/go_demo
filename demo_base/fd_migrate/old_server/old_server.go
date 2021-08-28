package main

import (
	"net"
	"syscall"
)

func main() {
	tcpAddr := net.TCPAddr{Port: 8001}
	tcpLn, err := net.ListenTCP("tcp4", &tcpAddr)
	if err != nil {
		panic(err)
	}
	f, _ := tcpLn.File()
	fdNum := f.Fd()
	data := syscall.UnixRights(int(fdNum))
	// 与新版本sidecar通过Unix Domain Socket建立链接
	raddr, _ := net.ResolveUnixAddr("unix", "/dev/shm/migrate.sock")
	uds, _ := net.DialUnix("unix", nil, raddr)
	// 通过UDS，发送ListenFD到新版本sidecar容器
	_, _, _ = uds.WriteMsgUnix(nil, data, nil)
	// 停止接收新的request，并且开始排水阶段
	tcpLn.Close()
}
