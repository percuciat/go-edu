package service

import (
	"sync"
)

type CustomType int

type PersonService interface {
	GetName() string
	GetAge() int
	ChangeAge(newAge int) bool
	CreateMap() map[string]int
}

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) CreateMap() map[string]int {
	return make(map[string]int)
}

func (p Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}

func (p *Person) ChangeAge(newAge int) bool {
	p.Age = newAge
	return true
}

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Service interface {
	GetItems() []Item
	GetItem(id int) (*Item, error)
	CreateItem(item Item) Item
	UpdateItem(id int, item Item) (*Item, error)
	DeleteItem(id int) error
}

type service struct {
	items  map[int]Item
	nextID int
	mutex  sync.Mutex
}

func New() Service {
	return &service{
		items:  make(map[int]Item),
		nextID: 1,
	}
}

func (s *service) GetItems() []Item {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	items := make([]Item, 0, len(s.items))
	for _, item := range s.items {
		items = append(items, item)
	}
	return items
}

func (s *service) GetItem(id int) (*Item, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if item, exists := s.items[id]; exists {
		return &item, nil
	}
	return nil, nil
}

func (s *service) CreateItem(item Item) Item {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	item.ID = s.nextID
	s.nextID++
	s.items[item.ID] = item
	return item
}

func (s *service) UpdateItem(id int, item Item) (*Item, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.items[id]; !exists {
		return nil, nil
	}

	item.ID = id
	s.items[id] = item
	return &item, nil
}

func (s *service) DeleteItem(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.items[id]; !exists {
		return nil
	}

	delete(s.items, id)
	return nil
}
