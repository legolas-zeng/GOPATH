package main
import (
	"fmt"
	"net/http"
	"log"
)
func main() {
	log.Print("启动服务...")
	http.HandleFunc("/home", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello! Your request was processed.")
	},
	)
	log.Print("服务启动完成.....")
	log.Fatal(http.ListenAndServe(":8000", nil))
}


//打开 http://127.0.0.1:8000/home