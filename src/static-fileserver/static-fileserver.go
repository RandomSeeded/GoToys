package main

import (
	// "github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	panic(http.ListenAndServe(":9000", http.FileServer(http.Dir("/Users/nate/Projects/go/public/spotifree")))) // WHAT DOES PANIC DO?

        // router := httprouter.New()
        // http.ListenAndServe(":9000", router)
}
