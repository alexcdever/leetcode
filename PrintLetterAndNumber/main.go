package main

import "sync"

func main() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	go func() {
		i := 0
		for {
			select {
			case <-number:
				print(i)
				i++
				letter <- true
			}
		}
	}()
	wait.Add(1)
	go func() {
		i := 'A'
		for {
			if i > 'Z' {
				wait.Done()
				return
			}
			select {
			case <-letter:
				print(string(i) + "\n")
				i++
				number <- true
			}
		}
	}()
	number <- true
	wait.Wait()
}
