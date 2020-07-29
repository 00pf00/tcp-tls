package main

import (
	"crypto/tls"
	"fmt"
	"reflect"
)

func main() {
	serverCert, err := tls.LoadX509KeyPair("/Users/feeli/workspace/tcp-tls/src/github.com/00pf00/tcp-tls/conf/server/server.crt", "/Users/feeli/workspace/tcp-tls/src/github.com/00pf00/tcp-tls/conf/server/server.key")
	if err != nil {
		fmt.Printf("load cert failed err = %v", err)
		return
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		//InsecureSkipVerify:true,
	}
	ln, err := tls.Listen("tcp", "0.0.0.0:9000", config)

	for true {
		conn, connerr := ln.Accept()
		if connerr != nil {
			fmt.Printf("accept failed err = %v", connerr)
			continue
		}
		fmt.Println(reflect.TypeOf(conn))
		tconn := conn.(*tls.Conn)

		fmt.Println("start")
		fmt.Println(tconn.ConnectionState().ServerName)
		fmt.Println("end")
		rb := make([]byte, 1)
		_, rerr := tconn.Read(rb)
		if rerr != nil {
			fmt.Printf("read failed err = %v", rerr)
		}
		fmt.Println(string(rb))
		fmt.Println(tconn.ConnectionState().ServerName)
	}
}
