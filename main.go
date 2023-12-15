package main

import (
	"fmt"
	"net/http"
)

// Структура альбомов
// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// TODO: Тестовые данные
// var albums = []album{
// 	{ID: "1", Title: "IGOR", Artist: "Tyler, The creator", Price: 3.99},
// 	{ID: "2", Title: "Ethernal blue", Artist: "Spiritbox", Price: 2.5},
// 	{ID: "3", Title: "Madvillain", Artist: "MF DOOM", Price: 5},
// }

func root(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "You can use endpoints such '/hello' and '/headers' for test")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello!\n")
}

// This handler does something a little more sophisticated by reading all the HTTP request headers and echoing them into the response body.
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func apiAlbum(w http.ResponseWriter, req *http.Request) {
	rId := req.URL.Path
	var retData = []album{}
	if len(rId) > 0 {
		for _, a := range albums {
			if a.ID == rId {
				retData = append(retData, a)
			}
		}
	} else {
		retData = albums
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(retData)
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":80", nil)
}
