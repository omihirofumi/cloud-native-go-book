package main

import (
	"fmt"
	"log"
	"os"
	"plugin"
)

type Sayer interface {
	Says() string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: run ./chap08/go-plugin/main/main.go animal")
	}

	name := os.Args[1]
	module := fmt.Sprintf("./chap08/go-plugin/%s/%s.so", name, name)

	p, err := plugin.Open(module)
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := p.Lookup("Animal")
	if err != nil {
		log.Fatal(err)
	}

	animal, ok := symbol.(Sayer)
	if !ok {
		log.Fatal("that's not a Sayer")
	}

	fmt.Printf("A %s says: %q\n", name, animal.Says())
}
