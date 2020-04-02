package main

import (
	"fmt"
	"reflect"
)

type Hoge struct {
	Name     string `elem:"n"`
	Email    string `elem:"e"`
	Password string `elem:"p" secret:"true"`
}

func main() {
	h := &Hoge{Name: "aaa", Email: "bbb", Password: "mysecret"}
	fmt.Println(h)
	v := reflect.ValueOf(h).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("secret") == "true" {
			fmt.Printf("Name=%s, Value=%s, tag(secret)=%v\n", field.Name, v.Field(i), true)
			v.Field(i).SetString("overwritten!")
		}
	}
	fmt.Println(h)
}
