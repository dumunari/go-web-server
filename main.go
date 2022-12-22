package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Go Web Server Started!!!!")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var req string

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	req = req + fmt.Sprintf("URL.Path = %q\n", r.URL.Path)

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		req = req + fmt.Sprintf("Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	req = req + fmt.Sprintf("Host = %q\n", r.Host)

	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		req = req + fmt.Sprintf("Form[%q] = %q\n", k, v)
	}

	hash := sha256.New()
	hash.Write([]byte(req))
	sha256_hash := hex.EncodeToString(hash.Sum(nil))

	fmt.Fprintf(w, "Request Hash: %s\n", sha256_hash)
}
