package aesample

import (
	"fmt"
	"net/http"

	"github.com/mzimmerman/aesample/subpkga"
	"github.com/mzimmerman/aesample/util"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/subpkga", subpkga.Root)
	util.Log("aesample init()")
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, you've reached aesample, the root package does nothing")
}
