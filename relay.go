package main

import "bufio"
import "bytes"
import "encoding/json"
import "flag"
import "fmt"
import "net/http"
import "os"
import "time"


type Message struct {
    Timestamp time.Time `json:"timestamp"`
    Level     string    `json:"level"`
    Sublevel  string    `json:"sublevel"`
    Message   string    `json:"message"`
}


func post_message(endpoint *string, level *string, sublevel *string, message string){
    m := Message{time.Now(), *level, *sublevel, message}
    b, err := json.Marshal(m)
    if err != nil {
        panic(err)
    }
    resp, err := http.Post(*endpoint, "application/json", bytes.NewBuffer(b))

    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()
    return
}


// The simplest use of a Scanner, to read standard input as a set of lines.

func main() {
    level := flag.String("level", "INFO", "level")
    sublevel := flag.String("sublevel", "user", "sublevel")
    endpoint := flag.String("endpoint", "http://127.0.0.1:9000", "endpoint")
    flag.Parse()

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        post_message(endpoint, level, sublevel, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}
