package main
import (
 "fmt"
 "net/http"
 "os"
)
func main() {
 var PORT string
 if PORT = os.Getenv("PORT"); PORT == "" {
  PORT = "5000"
 }
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi Anatoliy, could you buy me something...\n")
  fmt.Fprintf(w, "App port is: %s and URL path is: %s\n", PORT, r.URL.Path)
  fmt.Println("Anatoliy, you should buy me a beer!!!")
 })
 http.ListenAndServe(":" + PORT, nil)
}