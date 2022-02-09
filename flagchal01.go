/* Alta3 Research | RZFeeser
   CHALLEGE 01 - Reading and displaying "n" lines of a file */


package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
)

func main() {

    var count int
    flag.IntVar(&count, "n", 5, "number of lines to read from the file")
    flag.Parse()

    var in io.Reader
    // if flag.Arg(0) is "not" empty, it means there is a trailing arg
    // this trailing arg is the name of the file we need to scan
    if filename := flag.Arg(0); filename != "" {
        f, err := os.Open(filename)
        
        // error handling
        if err != nil {
            fmt.Println("error opening file: err:", err)
            os.Exit(1)
        }
        defer f.Close()   // still close the file if an error occurs

        in = f
    } else {
        // if the user did NOT pass a filename as an argument 
        // fmt.Println("You forgot to provide data to parse!")
        in = os.Stdin     // then we read from standard input
    }

    // read lines from the io.Reader var in
    buf := bufio.NewScanner(in)

    // iterate up to the value of count user a "for loop" 
    for i := 0; i < count; i++ {
        if !buf.Scan() {            // if scanning produces a false
            break                   // because we were asked to scan more lines
        }                           // than actually exist 
        fmt.Println(buf.Text())
    }

    if err := buf.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error reading: err:", err)
    }
}

