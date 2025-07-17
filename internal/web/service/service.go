package service

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

var (
	items  = make(map[int]Item)
	nextID = 1
	mutex  = &sync.Mutex{}
)

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mutex.Lock()
	newItem.ID = nextID
	nextID++
	items[newItem.ID] = newItem
	mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	itemList := make([]Item, 0, len(items))
	for _, item := range items {
		itemList = append(itemList, item)
	}
	json.NewEncoder(w).Encode(itemList)
}

func getItem(w http.ResponseWriter, r *http.Request, id int) {
	mutex.Lock()
	item, exists := items[id]
	mutex.Unlock()

	if !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request, id int) {
	var updatedItem Item
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mutex.Lock()
	_, exists := items[id]
	if !exists {
		mutex.Unlock()
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	updatedItem.ID = id
	items[id] = updatedItem
	mutex.Unlock()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedItem)
}

func deleteItem(w http.ResponseWriter, r *http.Request, id int) {
	mutex.Lock()
	_, exists := items[id]
	if !exists {
		mutex.Unlock()
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	delete(items, id)
	mutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
