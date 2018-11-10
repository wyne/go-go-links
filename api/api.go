package main

import (
  "encoding/json"
  "io/ioutil"
  "fmt"
  "net/http"
  "github.com/go-redis/redis"
)

type GoLink struct {
  Key string
  Url string
}

func main() {
  redis := redis.NewClient(&redis.Options{
    Addr:     "redis:6379",
    Password: "",
    DB:       0,
  })

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

    switch r.Method {
    case "PUT":
      body, err := ioutil.ReadAll(r.Body)
      if err != nil {
        fmt.Printf("Error reading body: %v", err)
        http.Error(w, "can't read body", http.StatusBadRequest)
        return
      }

      var goLink GoLink
      json.Unmarshal([]byte(body), &goLink)

      fmt.Printf("Creating %s -> %s\n", r.URL.Path, goLink.Url)
      seterr := redis.Set(r.URL.Path, goLink.Url, 0).Err()
      if seterr != nil {
        panic(seterr)
      }
    default:
      count, err := redis.Incr("count:" + r.URL.Path).Result()
      if err != nil {
        panic(err)
      }

      val, geterr := redis.Get(r.URL.Path).Result()
      if geterr != nil {
        panic(geterr)
      }

      fmt.Printf("Hit %s -> %s (%d)\n", r.URL.Path, val, count)
      http.Redirect(w, r, val, 301)
    }
  })

  http.ListenAndServe(":8080", nil)
}
