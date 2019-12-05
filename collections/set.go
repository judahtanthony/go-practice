package main

import (
	"fmt"

	"github.com/golang-collections/collections/set"
)

type MySet map[interface{}]bool

func (s *MySet) Insert(key interface{}) {
	(*s)[key] = true
}

func (s *MySet) Remove(key interface{}) {
	delete((*s), key)
}

func (s *MySet) Has(key interface{}) bool {
	return (*s)[key]
}

type MyCounter map[interface{}]int

func (c *MyCounter) Insert(key interface{}) {
	(*c)[key] += 1
}
func (c *MyCounter) Has(key interface{}) bool {
	_, ok := (*c)[key]
	return ok
}
func (c *MyCounter) Count(key interface{}) (int, bool) {
	val, ok := (*c)[key]
	return val, ok
}

func main() {
	fmt.Println("Set ======")
	s := set.New()
	s.Insert("a")
	s.Insert("b")
	s.Insert("a")
	fmt.Println(s.Has("a"))
	fmt.Println(s.Has("b"))
	fmt.Println(s.Has("c"))

	// Exercise MySet
	fmt.Println("MySet ======")
	ms := make(MySet)
	ms.Insert("a")
	ms.Insert("b")
	ms.Insert("a")
	fmt.Println(ms.Has("a"))
	fmt.Println(ms.Has("b"))
	fmt.Println(ms.Has("c"))

	// Exercise MyCounter
	fmt.Println("MyCounter ======")
	mc := make(MyCounter)
	mc.Insert("a")
	mc.Insert("b")
	mc.Insert("a")
	fmt.Println(mc.Has("a"))
	fmt.Println(mc.Has("b"))
	fmt.Println(mc.Has("c"))
	fmt.Println(mc.Count("a"))
	fmt.Println(mc.Count("b"))
	fmt.Println(mc.Count("c"))
}
