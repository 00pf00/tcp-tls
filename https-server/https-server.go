package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

type S struct {
}

func (s *S) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start")
	fmt.Println(r.TLS.ServerName)
	fmt.Println("end")

}
func main() {
	cert, err := tls.LoadX509KeyPair("/Users/feeli/workspace/tcp-tls/src/github.com/00pf00/tcp-tls/conf/server/server.crt", "/Users/feeli/workspace/tcp-tls/src/github.com/00pf00/tcp-tls/conf/server/server.key")
	if err != nil {
		fmt.Println("client load cert fail !")
	}
	sh := &S{}
	s := http.Server{
		Addr:    ":9000",
		Handler: sh,
		TLSConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		},
	}
	errs := s.ListenAndServeTLS("", "")
	fmt.Printf("server error = %v\n", errs)
}
