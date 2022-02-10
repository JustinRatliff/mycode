/* Alta3 Research | RZFeeser
   Channels - No channels, using a mutex  */

package main

import (
   "os"
   "fmt"
   "runtime"
   "strconv"
   "sync"
)

var accountBalance = 0    // balance for shared bank account
var mutex = &sync.Mutex{} // mutual-exclusion lock

// critical-section code with explicit locking/unlocking
func increaseBalance(amt int) {
   mutex.Lock()
   accountBalance += amt  // two operations: update and assignment
   mutex.Unlock()
}

func reportAndExit(msg string) {
   fmt.Println(msg)
   os.Exit(-1) // all 1s in binary
}

func main() {
   if len(os.Args) < 2 {
      reportAndExit("\nUsage: go run channels02.go <number of updates per thread>")
   }
   iterations, err := strconv.Atoi(os.Args[1])
   if err != nil {
      reportAndExit("Bad command-line argument: " + os.Args[1]);
   }

   var wg sync.WaitGroup  // wait group to ensure goroutine coordination

   // profiting increments the balance
   wg.Add(1)           // increment WaitGroup counter
   go func() {
      defer wg.Done()  // invoke Done on the WaitGroup when finished
      for i := 0; i < iterations ; i++ {
         increaseBalance(1)
         fmt.Println("Adding")
         runtime.Gosched()  // yield to another goroutine
         // using runtime.Gosched() will "yield" back to the main() thread
         // and the increaseBalance(-1) function will run
      }
   }()

   // spendy decrements the balance
   wg.Add(1)           // increment WaitGroup counter
   go func() {
      defer wg.Done()
      for i := 0; i < iterations; i++ {
         increaseBalance(-1)
         fmt.Println("Removing")
         runtime.Gosched()  // be nice--yield
      }
   }()

   wg.Wait()  // await completion of profiting and spendy
   fmt.Println("Final balance: ", accountBalance)  // confirm final balance is zero
}

