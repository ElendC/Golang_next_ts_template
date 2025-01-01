package main

import (
	"fmt"
	//"crypto/tls"
	"log"
	"net/http"
)

func ServerSetup() {

	port := ":443"

	fmt.Println("Backend Server Starting...")
	fmt.Println("Visit at: ")
	fmt.Printf("https://localhost%s", port)

	// TODO: For enhanced security etc., configure TLS settings.
	// Ref:
	// - tls.LoadX509KeyPair: Loads the certificate and private key.
	// - tls.Config: https://pkg.go.dev/crypto/tls#Config (for advanced TLS settings).
	// Example: Use http.Server with Addr, TLSConfig, and other configurations for production: https://pkg.go.dev/net/http#Server

	// TODO: Change the certFile and keyFile for production.
	// nil as handler because we have multiple handlers. the http.HandleFunc allows server to use http.DefaultServerMux,
	// which automatically routes requests to the appropriate handler
	log.Fatal(http.ListenAndServeTLS(":443", "certificates/localhost.pem", "certificates/localhost-key.pem", nil))

}
