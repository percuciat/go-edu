package gorutines

import (
	"fmt"
	"sync"
	"time"
)

type ActionsService interface {
	PrintNumbers(wg *sync.WaitGroup)
}

// Реализация интерфейса Printer
type actions struct{}

// NewConsolePrinter создаёт новый экземпляр consolePrinter
func NewActions() ActionsService {
	return &actions{}
}

func (a *actions) PrintNumbers(wg *sync.WaitGroup) {
  // добавляем wg чтобы завершить выполнение горутины
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}
