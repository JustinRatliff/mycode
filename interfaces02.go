/* RZFeeser | Alta3 Research
   Interfaces - Getting at interfaces */

package main

import "fmt"

type animal interface {
    breathe()
    walk()
}

type tiger struct {
     age int
}

func (l tiger) breathe() {
    fmt.Println("tiger breathes")
}

func (l tiger) walk() {
    fmt.Println("tiger walk")
}

type giraffe struct {
     age int
}

func (l giraffe) breathe() {
    fmt.Println("giraffe breathes")
}

func (l giraffe) walk() {
    fmt.Println("giraffe walk")
}

func main() {
    l := tiger{age: 10}
    callBreathe(l)
    callWalk(l)

    d := giraffe{age: 5}
    callBreathe(d)
    callWalk(d)
}

func callBreathe(a animal) {
    a.breathe()
}

func callWalk(a animal) {
    a.walk()
}

