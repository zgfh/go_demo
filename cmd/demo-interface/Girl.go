package main

import "log"

type Girl struct {
	Name string
}

func NewGirl(name string) *Girl {
	return &Girl{Name: name}
}
func (g *Girl) Say(word string) {
	log.Printf("girl %s say: %s", g.Name, word)
}
