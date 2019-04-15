package main

import (
	"github.com/4ndr3ye/DemoServer/handler"
	"github.com/4ndr3ye/DemoServer/middleware"
	"github.com/4ndr3ye/DemoServer/model"
	"crypto/tls"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	model.Connect(model.Databases)

	fmt.Println("The server is running on port 8443")

	r := mux.NewRouter()

	r.HandleFunc("/", handler.Base).Methods("GET")
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/login", handler.GetLogin).Methods("GET")
	r.HandleFunc("/custumers", middleware.AuthMiddleware(handler.GetCustumers)).Methods("GET")
	r.HandleFunc("/addcustumer", middleware.AuthMiddleware(handler.GetAddCustumers)).Methods("GET")
	r.HandleFunc("/addcustumer", middleware.AuthMiddleware(handler.SubmitCustumer)).Methods("POST")
	r.HandleFunc("/alive", middleware.AuthMiddleware(handler.GetCheckAlive)).Methods("GET", "POST")
	r.HandleFunc("/films", middleware.AuthMiddleware(handler.GetFilms)).Methods("GET")
	r.HandleFunc("/secret", middleware.AuthMiddleware(handler.Welcome)).Methods("GET")
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         ":8443",
		Handler:      r,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(srv.ListenAndServeTLS("tls.crt", "tls.key"))
}
