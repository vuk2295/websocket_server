package main

import (
	"chat/api"
	"chat/user"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
)

var rdb *redis.Client

func init() {
	rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	defer rdb.Close()
	rdb.SAdd(user.ChannelsKey, "general", "random")
}

func main() {

	rdb = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})

	r := mux.NewRouter()

	r.Path("/chat").Methods("GET").HandlerFunc(api.H(rdb, api.ChatWebSocketHandler))
	r.Path("/user/{user}/channels").Methods("GET").HandlerFunc(api.H(rdb, api.UserChannelsHandler))
	r.Path("/users").Methods("GET").HandlerFunc(api.H(rdb, api.UsersHandler))

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}
	fmt.Println("socket started on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}
