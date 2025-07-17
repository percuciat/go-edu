package storage

import (
	"errors"
)

type Employee struct {
	Id     int
	Name   string
	Age    int
	Salary int
}

type Storage interface {
	Get(id int) (Employee, error)
	Insert(e Employee) error
	Delete(id int) error
}

type EmployeeStorage struct {
	data map[int]Employee
}

func NewEmployeeStorage() *EmployeeStorage {
	return &EmployeeStorage{
		data: make(map[int]Employee),
	}
}

func (eS *EmployeeStorage) Get(id int) (Employee, error) {
	e, exists := eS.data[id]
	if !exists {
		return Employee{}, errors.New("employee not found")
	}
	return e, nil
}

func (eS *EmployeeStorage) Insert(e Employee) error {
	eS.data[e.Id] = e
	return nil
}

func (eS *EmployeeStorage) Delete(id int) error {
	delete(eS.data, id)
	return nil
}
