package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	clientCert, err := tls.LoadX509KeyPair("/Users/feeli/workspace/tcp-tls/src/github.com/00pf00/tcp-tls/conf/client/client.crt", "/Users/feeli/workspace/tcp-tls/src/github.com/00pf00/tcp-tls/conf/client/client.key")
	if err != nil {
		fmt.Printf("load cert failed err = %v", err)
		return
	}
	config := &tls.Config{
		Certificates:       []tls.Certificate{clientCert},
		InsecureSkipVerify: true,
		//ServerName:         "tunnel.cloud.com",
	}
	conn, err := tls.Dial("tcp", "tunnel.cloud.com:9000", config)
	if err != nil {
		fmt.Printf("connet failed err = %v", err)
		return
	}
	//req, _ := http.NewRequest(http.MethodGet,"/a",nil)
	//req.Write(conn)
	conn.Write([]byte{'a'})
	fmt.Println("end")

}
