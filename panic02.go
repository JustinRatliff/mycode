package main

import "os"

func main() {

    _, err := os.Create("/carolDanvers/msmarvel")
    
    // if we have an unexpected error
    if err != nil {
        panic(err)  // fast fail and dump the error to the screen
    }
}

