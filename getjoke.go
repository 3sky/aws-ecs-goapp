//bash.org.pl/random/

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

//GetJoke gets jokes
func GetJoke(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(DownloadJoke())
	// s :=
	// json.NewEncoder(w).Encode(s.SetEscapeHTML(false))
}

//SayHello seyhello
func SayHello(w http.ResponseWriter, r *http.Request) {
	hmgs := "Hello Message"
	json.NewEncoder(w).Encode(hmgs)
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
	router.HandleFunc("/hello", SayHello).Methods("GET")
	router.HandleFunc("/joke", GetJoke).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
