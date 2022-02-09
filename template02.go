/* Alta3 Research | RZFeeser
   Template - HTML template  */

package main


import (
    "html/template"
    "os"
)


type Todo struct {
    Title string
    Done  bool
}


type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}


func main() {

    // check our template for potential errors with Must
    tmpl := template.Must(template.ParseFiles("index.html"))
    
        data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Laundry", Done: false},
                {Title: "Pull weeds in the garden", Done: true},
                {Title: "Sweep the stairs", Done: true},
            },
        }
        tmpl.Execute(os.Stdout, data)
}

