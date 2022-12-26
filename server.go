package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

func handleUpgrade(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Connection") != "Upgrade" || r.Header.Get("Upgrade") != "MyProtocol" {
		w.WriteHeader(400)
		return
	}
	fmt.Println("Upgrade to MyProtocol")

	hijacker := w.(http.Hijacker)
	conn, readWriter, err := hijacker.Hijack()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	response := http.Response{
		StatusCode: 101,
		Header:     make(http.Header),
	}
	response.Header.Set("Upgrade", "MyProtocol")
	response.Header.Set("Connection", "Upgrade")
	response.Write(conn)

	for i := 1; i <= 10; i++ {
		fmt.Fprintf(readWriter, "%d\n", i)
		fmt.Println("send ->", i)
		readWriter.Flush()
		recv, err := readWriter.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("receive <- %s", string(recv))
		time.Sleep(500 * time.Millisecond)
	}
}

func handleChunkedResponse(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("Exptected http.ResponseWriter to be an http.Flusher")
	}
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(w, "Chunk #%d\n", i)
		flusher.Flush()
		time.Sleep(500 * time.Millisecond)
	}
	flusher.Flush()
}

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return

	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>/n")
}

type Calculator int
type Args struct {
	A, B int
}

func (c *Calculator) Multiply(args Args, result *int) error {
	log.Printf("Multiply called: %d, %d\n", args.A, args.B)
	*result = args.A * args.B
	return nil
}

func main() {
	Calculator := new(Calculator)
	server := rpc.NewServer()
	server.Register(Calculator)
	http.Handle(rpc.DefaultRPCPath, server)
	log.Println("start http listening :18888")
	listener, err := net.Listen("tcp", ":18888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

// func mainWithUpgrade() {
// 	var httpServer http.Server
// 	http.HandleFunc("/", handler)
// 	http.HandleFunc("/upgrade", handleUpgrade)
// 	http.HandleFunc("/chunked", handleChunkedResponse)
// 	log.Println("start http listening :18888")
// 	httpServer.Addr = ":18888"
// 	log.Println(httpServer.ListenAndServe())
// }
