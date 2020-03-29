package main

import (
	"fmt"
	"log"
	"time"
)

type Handler func(string) error

func Recover(h Handler) Handler {
	return func(arg string) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recoved from: %v", r)
			}
		}()

		return h(arg)
	}
}

func Async(h Handler) Handler {
	return func(arg string) error {
		go h(arg)
		return nil
	}
}

func Log(h Handler) Handler {
	return func(arg string) error {
		err := h(arg)
		log.Printf("called f with %s, returns %v", arg, err)
		return err
	}
}

func hello(name string) error {
	fmt.Printf("Hello, %s\n", name)
	panic("Ouch")
}

var Hello = Async(Log(Recover(hello)))

func main() {
	Hello("Beoran")
	time.Sleep(time.Second)
}
