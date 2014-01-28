package subpkgb

import (
	"aesample/util"
	"fmt"
	"net/http"
)

func init() {
	util.Log("init() in subpkgb")
	http.HandleFunc("/subpkgb", Root)
}

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, you've reached aesample/subpkgb")
}
