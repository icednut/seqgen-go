package main

import (
    "fmt"
    "net/http"
)

func seqGenerator(sequenceReceiveChannel <-chan int, sequenceGenerateChannel chan<- int) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        sequenceGenerateChannel <- 1
        fmt.Fprintf(w, "%d", <-sequenceReceiveChannel)
    }
}

func runServer() {
    sequenceReceiveChannel := make(chan int, 1)
    sequenceGenerateChannel := make(chan int, 1)

    go func() {
        sequence := 0
        for {
            sequencePiece := <-sequenceGenerateChannel
            sequence = sequence + sequencePiece
            //fmt.Println("seqNo: ", sequence)
            sequenceReceiveChannel <- sequence
        }
    }()

    http.HandleFunc("/", seqGenerator(sequenceReceiveChannel, sequenceGenerateChannel))
    fmt.Println("Waiting for request... (localhost:8080)")
    http.ListenAndServe(":8080", nil)
}

func main() {
    runServer()
}
