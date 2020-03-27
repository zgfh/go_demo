package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

/*
参考： https://golang.org/pkg/encoding/json/
json 格式 https://golang.org/pkg/encoding/json/#Marshal
*/
type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}

	return nil
}

func (a Animal) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}

	return json.Marshal(s)
}

func main() {
	gopher := "gopher"
	strB, _ := json.Marshal(gopher)
	log.Printf("gopher byte %s", string(strB))

	var animal *Animal
	if err := json.Unmarshal(strB, &animal); err != nil {
		log.Fatal(err)
	}
	log.Printf("animal %s", animal)

	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	var zoo []*Animal
	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
	}
	log.Printf("zoo %s", zoo)

	census := make(map[Animal]int)
	for _, animal := range zoo {
		census[*animal] += 1
	}

	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras:  %d\n* Unknown: %d\n",
		census[Gopher], census[Zebra], census[Unknown])

}
