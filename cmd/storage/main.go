package main

import (
	"fmt"
	"strconv"
	"goProject/internal/storage"
)

type a interface{}

func main() {
	var a a
	s := storage.NewEmployeeStorage()
	createMockStorage(s)
	v, ok := a.(string)
	fmt.Print(v, ok)
	//fmt.Print(a)
	/* for k, v := range s.data {
	   fmt.Printf("id: %d, name: %s, age: %d salary: %d\n", k, v.name, v.age, v.salary)
	 }  */
}

func createMockStorage(s storage.Storage) {
	for i := 1; i <= 10; i++ {
		e := storage.Employee{Name: "Anton" + strconv.Itoa(i), Age: 20 + i, Id: i, Salary: i + 100}
		s.Insert(e)
		//fmt.Printf("Number: %d, String: %s\n", 42, "example")
		//fmt.Printf("id: %d, name: %s, age: %d salary: %d\n", e.id, e.name, e.age, e.salary)
	}
}
