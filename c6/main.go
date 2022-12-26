package main

import (
	"log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:18888")
	if err != nil {
		panic(err)
	}
	var result int
	args := &Args{4, 5}
	err = client.Call("Calculator.Multiply", args, &result)
	if err != nil {
		panic(err)
	}
	log.Printf("4 x 5 = %d\n", result)
}

// func main() {

// 	dialer := &net.Dialer{
// 		Timeout:   30 * time.Second,
// 		KeepAlive: 30 * time.Second,
// 	}
// 	conn, err := dialer.Dial("tcp", "localhost:18888")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()

// 	request, err := http.NewRequest("GET", "http://localhost:18888/chunked", nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = request.Write(conn)
// 	if err != nil {
// 		panic(err)
// 	}

// 	reader := bufio.NewReader(conn)
// 	resp, err := http.ReadResponse(reader, request)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if resp.TransferEncoding[0] != "chunked" {
// 		panic("wrong transfer encoding")
// 	}
// 	for {
// 		sizeStr, err := reader.ReadBytes('\n')
// 		if err == io.EOF {
// 			break
// 		}

// 		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
// 		if size == 0 {
// 			break
// 		}
// 		if err != nil {
// 			panic(err)
// 		}
// 		line := make([]byte, int(size))
// 		reader.Read(line)
// 		reader.Discard(2)
// 		log.Println(line)
// 		log.Println(" ", string(line))
// 	}
// }

// func mainWithChunk() {

// 	dialer := &net.Dialer{
// 		Timeout:   30 * time.Second,
// 		KeepAlive: 30 * time.Second,
// 	}
// 	conn, err := dialer.Dial("tcp", "localhost:18888")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()
// 	reader := bufio.NewReader(conn)

// 	request, _ := http.NewRequest("GET", "http://localhost:18888/upgrade", nil)
// 	request.Header.Set("Connection", "Upgrade")
// 	request.Header.Set("Upgrade", "MyProtocol")
// 	err = request.Write(conn)
// 	if err != nil {
// 		panic(err)
// 	}

// 	resp, err := http.ReadResponse(reader, request)
// 	if err != nil {
// 		panic(err)
// 	}

// 	log.Println("Status:", resp.Status)
// 	log.Println("Headers:", resp.Header)

// 	counter := 10
// 	for {
// 		data, err := reader.ReadBytes('\n')
// 		if err == io.EOF {
// 			break
// 		}
// 		fmt.Println("receive <-", string(bytes.TrimSpace(data)))
// 		fmt.Fprintf(conn, "%d\n", counter)
// 		fmt.Println("send ->", counter)
// 		counter--
// 	}
// }

// func mainWithCrt() {
// 	cert, err := tls.LoadX509KeyPair("../client.crt", "../client.key")
// 	if err != nil {
// 		panic(err)
// 	}

// 	client := &http.Client{
// 		Transport: &http.Transport{
// 			TLSClientConfig: &tls.Config{
// 				Certificates: []tls.Certificate{cert},
// 			},
// 		},
// 	}
// 	resp, err := client.Get("https://localhost:18443")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer resp.Body.Close()
// 	dump, err := httputil.DumpResponse(resp, true)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println(string(dump))
// }
