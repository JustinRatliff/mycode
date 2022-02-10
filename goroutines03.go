/* Alta3 Reseach | RZFeeser
   Goroutine - No Mutex     */
   
package main

import (
    "fmt"
    "sync"
)

var (
    mutex   sync.Mutex
    balance int
)

func init() {
    balance = 1000
}

func deposit(value int, wg *sync.WaitGroup) {
    // mutex.Lock()
    fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
    balance += value
    // mutex.Unlock()
    wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
    // mutex.Lock()
    fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
    balance -= value
    // mutex.Unlock()
    wg.Done()
}

func main() {
    fmt.Println("Go 'No Mutex' Example")

    var wg sync.WaitGroup
    wg.Add(3)
    go withdraw(750, &wg)   // if you pass a WaitGroup to a function, use a pointer
    go deposit(1000, &wg)
    go withdraw(500, &wg)
    wg.Wait()               // wait until we finish

    fmt.Printf("New Balance %d\n", balance)
}

