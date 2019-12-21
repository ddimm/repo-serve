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
		log.Fatal("couldn't reach api")
	}
	req.Header.Add("Authorization", "token "+os.Getenv("GITHUB_TOKEN"))
	resp, err := client.Do(req)
	fmt.Println(resp.Header.Get(("X-RateLimit-Remaining")))
	if err != nil {
		log.Fatal("couldn't reach api")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	fmt.Fprintf(*w, strBody)

}
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	fetchapi(&w)

}
func main() {
	s := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        http.HandlerFunc(handler),
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 50,
	}
	log.Fatal(s.ListenAndServe())
}
