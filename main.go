package main

 import (
         "fmt"
         "github.com/gorilla/mux"
         "net/http"
 )
// hello world function - who doesn't love Slick Rick?
 func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
         w.Write([]byte("Hello Young World! - Slick Rick"))
         // fmt.Fprintln(w,"trap or die")
         // fmt.Fprintln(w,"trap or die")
 }
// a service that read the vars provide via the URL to say hello
 func Sup(w http.ResponseWriter, r *http.Request) {
         name := mux.Vars(r)["name"]
         w.Write([]byte(fmt.Sprintf("Hello %s !", name)))
 }

// a service that will able a URL like
//    http://website:8080/person/Macolmn/leader/25
func ProcessPathVariables(w http.ResponseWriter, r *http.Request) {

         // break down the variables for easier assignment
         vars := mux.Vars(r)
         name := vars["name"]
         job := vars["job"]
         age := vars["age"]
         w.Write([]byte(fmt.Sprintf("Name is %s ", name)))
         w.Write([]byte(fmt.Sprintf("Title is %s ", job)))
         //fmt.Fprintf(w,"Age is %s",age)
         w.Write([]byte(fmt.Sprintf("Age is %s ", age)))
 }
// customize the 404 when you try to route to a path no listed

func notFound(w http.ResponseWriter, r *http.Request) {
          fmt.Fprintln(w, "404 Error homie - get right or you will never make it")

}
 func main() {
         mx := mux.NewRouter()

         mx.HandleFunc("/", SayHelloWorld)
         mx.HandleFunc("/{name}", Sup)

         mx.HandleFunc("/person/{name:[A-Za-z]+}/{job}/{age:[0-9]+}", ProcessPathVariables)

         mx.NotFoundHandler = http.HandlerFunc(notFound)

         http.ListenAndServe(":8732", mx)
 }