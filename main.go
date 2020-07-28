package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"../controllers"
	"github.com/julienschmidt/httprouter"
)

// https://pace.dev/blog/2020/02/12/why-you-shouldnt-use-func-main-in-golang-by-mat-ryer

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout io.Writer) error {

	// Init setup
	zips := controllers.ZipInitialize()
	// Launch server
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/", index)
	r.GET("/user/:id", uc.GetUser)
	http.ListenAndServe("localhost:8080", r)

	return nil
}

func main() {

}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html><html lang="en"><head><title>index</title></head>
	<body>
	1.0
	</body></html>`

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}
