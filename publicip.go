package main

import (
  "net/http"
  "encoding/json"
  "fmt"
  "log"
)

type Body struct {
  Ip string
}

func main() {
  resp, requestErr := http.Get("https://api.ipify.org?format=json")
  if requestErr != nil {
    log.Fatalln(requestErr)
  }

  defer resp.Body.Close()

  body := Body{}
  decodingErr := json.NewDecoder(resp.Body).Decode(&body)
  if decodingErr != nil {
    log.Fatalln(decodingErr)
  }

  fmt.Println(body.Ip)
}
