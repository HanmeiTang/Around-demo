package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Post struct {
	// `json:"user"` is for the json parsing of this User field.
	// Otherwise, by default it's 'User'.
	User     string   `json:"user"`
	Message  string   `json:"message"`
	Location Location `json:"location"`
	Url      string   `json:"url"`
	Type     string   `json:"type"`
	Face     float32  `json:"face"`
}

func main() {
	fmt.Println("started-service")
	http.HandleFunc("/post", handlerPost)
	// Fatal = print output and exit
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// Refer to "JSON and Go"
// If a user sends a HTTP request with a body as
// {
//   “user”: “jack”
//   “message”: “this is a message”
// }
// Then it will automatically construct a Post object named p
// and update its values to be p.User = “jack”
// and p.Message = “this  is a message”
func handlerPost(w http.ResponseWriter, r *http.Request) {
	// Parse from body of request to get a json object.
	fmt.Println("Received one post request")
	decoder := json.NewDecoder(r.Body)
	var p Post
	if err := decoder.Decode(&p); err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Post received: %s\n", p.Message)
}
