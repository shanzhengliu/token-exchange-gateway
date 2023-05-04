package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var client = &http.Client{
			Timeout: 5 * time.Second,
		}
		fmt.Println("request in the sidecar ----- handler token, and send the request to main")

		req, err := http.NewRequest(r.Method, "http://localhost:8081"+r.RequestURI, r.Body)

		req.Header.Add("Authorization", "hello")
		fmt.Println(fmt.Println("request in the sidecar ----- use new token"))
		if err != nil {
			fmt.Println(err)
		}
		rsps, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer rsps.Body.Close()
		body, _ := ioutil.ReadAll(rsps.Body)
		w.Header().Set("Succeed", "application/json")
		w.Write([]byte(body))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
