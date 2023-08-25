import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
	log.Println("Received request:", r.URL.Path)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the website")
	log.Println("Received request:", r.URL.Path)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server connected")
	log.Println("Received request:", r.URL.Path)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/about", aboutHandler)

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
