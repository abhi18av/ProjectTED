package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
)

func main() {
    url := "https://pi.tedcdn.com/r/pe.tedcdn.com/images/ted/6b6eb940bceab359ca676a9b486aae475c1df883_2880x1620.jpg?cb=20160511&quality=63&u=&w=512"
    // don't worry about errors
    response, e := http.Get(url)
    if e != nil {
        log.Fatal(e)
    }

    defer response.Body.Close()

    //open a file for writing
    file, err := os.Create("/home/abhi18av/ken.jpg")
    if err != nil {
        log.Fatal(err)
    }
    // Use io.Copy to just dump the response body to the file. This supports huge files
    _, err = io.Copy(file, response.Body)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()
    fmt.Println("Success!")
}