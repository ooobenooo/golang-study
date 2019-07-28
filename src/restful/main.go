package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/drone/routes"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are get user %s", uid)
}

func main() {
	mux := routes.New()
	mux.Get("/user/:uid", getUser)
	http.Handle("/", mux)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
}
