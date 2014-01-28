package subpkga

import (
	"aesample/util"
	"fmt"
	"net/http"
)

func init() {
	util.Log("subpkga init()")
}

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, you've reached aesample/subpkga")
}
