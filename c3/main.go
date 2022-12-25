package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	// file, err := os.Open("main.go")
	// if err != nil {
	// 	panic(err)
	// }

	reader := strings.NewReader("text1234")

	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		// 送信失敗
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
