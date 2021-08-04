package main
import ( 
    "fmt"
    "github.com/nats-io/nats.go"
)
func main() {
    fmt.Println("hello world")
    k := []byte{1, 2, 3}
    fmt.Println(k, len(k))

    // Connect to NATS
    nc, _ := nats.Connect("localhost:4222")

    // Create JetStream Context
    js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

    // Simple Stream Publisher
    js.Publish("test.subjects.hello", []byte("hello Stream test"))


}