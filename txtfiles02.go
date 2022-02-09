/* RZFeeser | Alta3 Research
   Creating text files          */    
     
package main
 
import (
    "bufio"
    "log"
    "os"
)
 
func main() {
    sampledata := []string{"West of House.",
        "You are standing in an open field west of a white house,",
        "with a boarded front door.",
        "There is a small mailbox here.",
    }
 
    file, err := os.OpenFile("ZorkOpening.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
 
    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }
 
    datawriter := bufio.NewWriter(file)
 
    for _, data := range sampledata {
        _, _ = datawriter.WriteString(data + "\n")
    }
 
    datawriter.Flush()
    file.Close()
}

