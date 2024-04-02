package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func towerOfHanoi(n int, source string, help string, dest string, wg *sync.WaitGroup) {
	defer wg.Done()

	if n == 1 {
		mutex.Lock()
		fmt.Printf("Move disk from %s to %s\n", source, dest)
		mutex.Unlock()
		return
	}

	newWg := &sync.WaitGroup{} // finally making updates in the waitgroups waiting list

	newWg.Add(1)
	go towerOfHanoi(n-1, source, dest, help, newWg)
	newWg.Add(1)
	go towerOfHanoi(n-1, help, source, dest, newWg)

	newWg.Wait()

	mutex.Lock()
	fmt.Printf("Move disk from %s to %s\n", source, dest)
	mutex.Unlock()
}

func main() {
	Time := time.Now()
	var disks int
	fmt.Print("Enter number of disks: ")
	fmt.Scan(&disks)
	var wg sync.WaitGroup
	wg.Add(1)
	go towerOfHanoi(disks, "A", "B", "C", &wg)

	wg.Wait()
	fmt.Println("Time taken: ", time.Since(Time))
}
