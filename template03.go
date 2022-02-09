/* Alta3 Research | RZFeeser
   Template - Using a Map to build a template  */

package main


import (
    "html/template"
    "os"
    "log"
    "io"
)

func parse(path string) {
    t, err := template.ParseFiles(path)
    if err != nil {
        log.Print(err)
        return
    }

    f, err := os.Create(path)
    if err != nil {
        log.Println("create file: ", err)
        return
    }

    // A sample config
    config := map[string]string{
        "textColor":      "#abcdef",
        "linkColorHover": "#ffaacc",
    }

    err = t.Execute(f, config)
    if err != nil {
        log.Print("execute: ", err)
        return
    }
    f.Close()
}

func main() {
    
    f, _ := os.Create("example.css")
    
    f.Write([]byte("The text color is {{.textColor}} and the link color is {{.linkColorHover}}"))
    
    f.Close()
    
    parse("example.css")
    
    f, _ = os.Open("example.css")
    io.Copy(os.Stdout, f)
    f.Close()
}

