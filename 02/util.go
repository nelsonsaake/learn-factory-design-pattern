package main

import (
	"fmt"
	"log"
)

func cout(v ...any) {
	fmt.Println(v...)
}

func die(v ...any) {
	log.Fatalln(v...)
}
