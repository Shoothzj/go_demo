package fd_migrate

import (
	"net"
	"net/http"
	"os"
	"syscall"
)

// https://mp.weixin.qq.com/s/tTipcU8MKxTpFurWNEUIww
func migrateFd() {
	// ignore exception
	// old sidecar migrate to new sidecar
	// tcpLn *net.TCPListener
	var tcpLn *net.TCPListener
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
	// 新版本sidecar 接收ListenFD，并且开始对外服务
	// 监听UDS
	addr, _ := net.ResolveUnixAddr("unix", "/dev/shm/migrate.sock")
	unixLn, _ := net.ListenUnix("unix", addr)
	conn, _ := unixLn.AcceptUnix()
	buf := make([]byte, 32)
	oob := make([]byte, 32)
	// 接收 ListenFD
	_, oobn, _, _, _ := conn.ReadMsgUnix(buf, oob)
	scms, _ := syscall.ParseSocketControlMessage(oob[:oobn])
	if len(scms) > 0 {
		// 解析FD，并转化为 *net.TCPListener
		fds, _ := syscall.ParseUnixRights(&scms[0])
		f := os.NewFile(uintptr(fds[0]), "")
		ln, _ := net.FileListener(f)
		tcpLn, _ := ln.(*net.TCPListener)
		http.Serve(tcpLn, nil)
	}

}
