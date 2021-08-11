
package main

import (
  "fmt"
  "bytes"
  "mime/multipart"
  "net/http"
  "io/ioutil"
  "github.com/nats-io/nats.go"
  "reflect"
)

func main() {

  nc, _ := nats.Connect("localhost:4222")  

  js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))
  sub, err := js.PullSubscribe("test.subjects.rocket_data", "MONITOR")
  if err != nil {
    fmt.Println(err)
    return
  }
  msgs, err := sub.Fetch(10)
  //fmt.Println(string(msgs))
  for i := 0; i < len(msgs); i++ {
    fmt.Println( *msgs[i] )
    fmt.Println(reflect.TypeOf(*msgs[i]))
  }
  //fmt.Println(*msgs)
  fmt.Println(reflect.TypeOf(msgs))
  //convertMsg(msgs)
  url := "https://api.spacexdata.com/v3/rockets/falcon9"
  method := "GET"

  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  err1 := writer.Close()
  if err1 != nil {
    fmt.Println(err)
    return
  }


  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Set("Content-Type", writer.FormDataContentType())
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))

  js.Publish("test.subjects.rocket_data", []byte(body))
}

// func convertMsg(m []*nats.Msg) {
//   fmt.Printf("Received a message: %s\n", string(m.Data))
// }