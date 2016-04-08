package main



import (

	"log"
	"fmt"

	"net/http"

)


const LEN_APIV1 = len("/api/v1/")

type ApiV1Handler struct {}

type ApiV1SomeAPIHandler struct {}

func (handler *ApiV1SomeAPIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("(APIV1SomeAPI handler) path = ", r.URL.Path)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[LEN_APIV1:])

}



func (handler *ApiV1Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("(APIV1 handler) path = ", r.URL.Path)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[LEN_APIV1:])

}


func handler(w http.ResponseWriter, r *http.Request) {

	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	log.Println("(default handler) path = ", r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Welcome to the home page!")
}



func main() {

	mux := http.NewServeMux()
	mux.Handle("/api/v1/someapi", &ApiV1SomeAPIHandler{})
	mux.Handle("/api/v1/", &ApiV1Handler{})
	mux.HandleFunc("/", handler)


	http.ListenAndServeTLS(":443", "server.pem", "server.key", mux)

}
