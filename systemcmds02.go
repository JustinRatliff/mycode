/* Alta3 Research | RZFeeser
   exec.Command("program", "argument1", "argument2", "argument3"...)
   Transform strings into all caps with linux 'tr' controlled from GO   */

package main

import (
    "bytes"
    "fmt"
    "log"
    "os/exec"
    "strings"
)

func main() {
    cmd := exec.Command("tr", "a-z", "A-Z")

    cmd.Stdin = strings.NewReader("hit the hyperdrive chewy")

    var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("translated phrase: %q\n", out.String())
}

