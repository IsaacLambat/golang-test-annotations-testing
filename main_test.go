package main

import (
	"log"
	"testing"
)

func TestRepro(t *testing.T) {
	log.Println("This is a test")
	log.Println("This is a test")
	log.Println("This is a test")
	log.Println("This is a test")
	log.Println("This is a test")
	log.Println("This is a test")
	log.Println("This is a test")
	log.Println("This is a test")
	log.Println("This is a test")
	t.FailNow()
}
