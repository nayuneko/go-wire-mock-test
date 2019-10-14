package main

import (
	"log"
)

func main() {
	s := InitializeService()
	if err := s.CreateEntry(1); err != nil {
		log.Fatal(err)
	}
}
