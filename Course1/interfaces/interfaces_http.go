package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// custom writer
type logWriter struct{}

func (logWriter) Write(byteSlice []byte) (int, error) {
	fmt.Println("No. of bytes written", len(byteSlice))

	return len(byteSlice), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// fmt.Println("Resp: ", resp)

	// byteSlice := make([]byte, 99999) // initialise with 99999 bytes of empty space
	// resp.Body.Read(byteSlice)
	// fmt.Println(string(byteSlice))

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}
