package main

import (
	"flag"
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"log"
)

type echoServer struct {
	gnet.BuiltinEventEngine

	eng       gnet.Engine
	addr      string
	multicore bool
}

//v1
//type echoServer struct {
//	*gnet.EventServer
//}
//
//func (es *echoServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
//	log.Printf("Echo server is listening on %s (multi-cores: %t, event-loops: %d)\n",
//		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
//	return
//}
//
//func (es *echoServer) React(packet []byte, c gnet.Conn) (out []byte, action gnet.Action) {
//	// Echo synchronously.
//	return packet, gnet.None
//
//	/*
//		// Echo asynchronously.
//		data := append([]byte{}, packet...)
//		go func() {
//			time.Sleep(time.Second)
//			c.AsyncWrite(data)
//		}()
//		return
//	*/
//}

func (es *echoServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Printf("Socket with addr: %s has been opened...\n", c.RemoteAddr().String())
	return
}

func (es *echoServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.eng = eng
	log.Printf("echo server with multi-core=%t is listening on %s\n", es.multicore, es.addr)
	return gnet.None
}

func (es *echoServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	log.Printf("buf=%s\n", string(buf))
	c.Write(buf)
	return gnet.None
}

// linux nc -v ip_address port
func main() {
	var port int
	var multicore bool

	// Example command: go run echo.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 9000, "--port 9000")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.Parse()
	echo := &echoServer{addr: fmt.Sprintf("tcp://:%d", port), multicore: multicore}
	log.Fatal(gnet.Run(echo, echo.addr, gnet.WithMulticore(multicore)))
}
