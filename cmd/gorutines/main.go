package main

import (
	"fmt"
	"goProject/internal/gorutines"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	p := gorutines.NewActions()
	wg.Add(1)
	go p.PrintNumbers(&wg)
	wg.Wait()
	fmt.Print("gorutines DONE")
}
