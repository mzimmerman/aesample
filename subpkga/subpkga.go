package subpkga

import (
	"fmt"
	"net/http"

	"github.com/mzimmerman/aesample/util"
)

func init() {
	util.Log("subpkga init()")
}

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, you've reached aesample/subpkga")
}
