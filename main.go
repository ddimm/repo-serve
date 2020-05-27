package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func fetchapi(w *http.ResponseWriter) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/users/"+os.Getenv("GITHUB_USER")+"/repos", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	req.Header.Add("Authorization", "token "+os.Getenv("GITHUB_TOKEN"))
	resp, err := client.Do(req)
	fmt.Println(resp.Header.Get(("X-RateLimit-Remaining")))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	fmt.Fprintf(*w, strBody)

}
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fetchapi(&w)

}
func main() {
	port, hasPort := os.LookupEnv("PORT")
	if !hasPort {
		port = "8080"
	}
	fmt.Println(port)
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        http.HandlerFunc(handler),
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 50,
	}

	log.Fatal(s.ListenAndServe())
}
