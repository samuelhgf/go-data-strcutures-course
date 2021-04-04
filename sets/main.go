package main

import (
	"errors"
	"fmt"
)

// Set - representation of a set data structure
type Set struct {
	Elements map[string]struct{}
}

func NewSet() *Set {
	var set Set
	set.Elements = make(map[string]struct{})

	return &set
}

// Add - adds an element to our Set
func (s *Set) Add(elem string) {
	s.Elements[elem] = struct{}{}
}

// Delete - removes an element from our set if it exists
func (s *Set) Delete (elem string) error {
	if _, exists := s.Elements[elem]; !exists {
		return errors.New("element not present in set")
	}
	delete(s.Elements, elem)
	return nil
}

// Contains - checks to see if an element exists within the set
func (s *Set) Contains(elem string) bool {
	_, exists := s.Elements[elem]
	return exists
}

// List - prints out all of the elements within our set
func (s *Set) List() {
	for k, _ := range s.Elements {
		fmt.Println(k)
	}
}

func main() {
	fmt.Println("Sets Tutorial")

	// Initialize a Set
	mySet := NewSet()

	// Populate the Set
	mySet.Add("Earth")
	mySet.Add("Venus")
	mySet.Add("Mars")

	// Add the same element to a Set to see that it will overwrite the existing one
	mySet.Add("Earth")

	// Remove an element from a set successfully
	if err := mySet.Delete("Venus"); err != nil {
		fmt.Println(err.Error())
	}

	// Try to remove a non existing element from a set and print error message
	if err := mySet.Delete("Not in set"); err != nil {
		fmt.Println(err.Error())
	}

	// List all elements from the Set
	mySet.List()

	// Verify if an element exists in the set and print if is true or false
	fmt.Println(mySet.Contains("Mars"))
}
