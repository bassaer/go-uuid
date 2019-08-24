package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id1 := uuid.New()
	fmt.Printf("id1 : %#v\n", id1.String())
	if id2, err := uuid.NewRandom(); err == nil {
		fmt.Printf("id2 : %#v\n", id2.String())
	}
}
