package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"time"
)

var newCache *cache.Cache

func init() {
	newCache = cache.New(1*time.Minute, 20*time.Minute)
	newCache.Set("foo", "Miles", cache.DefaultExpiration)
}

func main() {
	http.HandleFunc("/", getFromCache)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}

func getFromCache(w http.ResponseWriter, r *http.Request) {
	foo, found := newCache.Get("foo")
	if !found {
		newCache.Set("foo", "Bob", cache.DefaultExpiration)
		foo, _ = newCache.Get("foo")
	}
	fmt.Fprintf(w, "Hello "+foo.(string))
}
