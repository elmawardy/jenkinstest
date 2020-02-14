package main

import "fmt"
import "net/http"
import "log"

func main() {
  for i:=0;i<12;i++{
	fmt.Println("Allah Akbar")
  }
 
   http.HandleFunc("/", handler)
   log.Fatal(http.ListenAndServe(":8081", nil))


}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi I'm version 9")
}
