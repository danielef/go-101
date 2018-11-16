package main
import (
	"fmt"
	"net/http"
	"strings"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)

	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
