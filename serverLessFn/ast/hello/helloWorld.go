package hello

import (
	"net/http"
)

func ok() {
	println("ok")
}

// HelloWorld writes "Hello, World!" to the HTTP response.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL)
	ok()
	// fmt.Fprint(w, "Hello, World!\n")
}
