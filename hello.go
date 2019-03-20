package main

import (
	"fmt"
	"log"

	"github.com/EvanXzj/stringutil"
)

type Hello struct {
	name string
	age  int
	sex  bool
}

func main() {
	fmt.Println(stringutil.Reverse("!ooooooooG ,olleH"))

	h := Hello{
		"CH",
		11,
		false,
	}

	log.Println(h)
}
