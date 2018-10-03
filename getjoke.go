package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//GetJoke gets jokes
func GetJoke(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(DownloadJoke())

}

//SayHello seyhello
func SayHello(w http.ResponseWriter, r *http.Request) {
	hmgs := "Hello Message"
	json.NewEncoder(w).Encode(hmgs)
}

//Check will work on `/`
func Check(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	cmsg := fmt.Sprintf("It work at: %s", dt.String())
	json.NewEncoder(w).Encode(cmsg)
}

//DownloadJoke get from bash.org.pl/random
func DownloadJoke() string {

	var trimJoke string
	res, err := http.Get("http://bash.org.pl/random/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".quote").Each(func(i int, s *goquery.Selection) {
		joke := s.Text()
		trimJoke = strings.TrimSuffix(strings.TrimSpace(joke), "\n")

	})
	return trimJoke
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Check).Methods("GET")
	router.HandleFunc("/hello", SayHello).Methods("GET")
	router.HandleFunc("/joke", GetJoke).Methods("GET")
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	http.ListenAndServe(":8080", loggedRouter)
}
