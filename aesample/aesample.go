package aesample

import (
	"aesample/subpkga"
	"aesample/util"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/subpkga", subpkga.Root)
	util.Log("aesample init()")
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, you've reached aesample, the root package does nothing")
}
