package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	lsnr, err := net.Listen("tcp", ":")
	if err != nil {
		fmt.Println("Error listening:", err)

		os.Exit(1)
	}
	fmt.Println("Listening on:", lsnr.Addr())
	err = http.Serve(lsnr, nil)
	fmt.Println("Server exited with:", err)
	os.Exit(1)
}
