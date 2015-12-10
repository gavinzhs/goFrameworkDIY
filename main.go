package main

import (
	"log"
	"net/http"
)

func main() {
	print("start")

	http.Handle("/", http.FileServer(http.Dir("/Users/zhengsai/Downloads/")))
	http.Handle("/curl/", http.StripPrefix("/curl/", http.FileServer(http.Dir("/Users/zhengsai/Downloads/curl"))))
	http.Handle("/make/", http.StripPrefix("/make/", http.FileServer(http.Dir("/Users/zhengsai/Downloads/make"))))

	log.Fatalln(http.ListenAndServe(":9000", nil))
}
