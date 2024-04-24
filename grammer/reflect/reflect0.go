package main

import (
	"fmt"
	"reflect"
)

type Order struct {
	ordId      int    `json:"order_id" validate:"required"`
	customerId string `json:"customer_id" validate:"required"`
	callback   func() `json:"call_back" validate:"required"`
}

func reflectInfo(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	for i := 0; i < v.NumField(); i = i + 1 {
		fv := v.Field(i)
		ft := t.Field(i)
		tag := t.Field(i).Tag.Get("json")
		validate := t.Field(i).Tag.Get("validate")
		switch fv.Kind() {
		case reflect.String:
			fmt.Printf("The %d th %s types: %s, valuing: %s, struct tag: %v\n", i, ft.Name, "string", fv.String(), tag+" "+validate)
		case reflect.Int:
			fmt.Printf("The %d th %s types %s, valuing %d, struct tag: %v\n", i, ft.Name, "int", fv.Int(), tag+" "+validate)
		case reflect.Func:
			fmt.Printf("The %d th %s types %s, valuing %v, struct tag: %v\n", i, ft.Name, "func", fv.String(), tag+" "+validate)
		}
	}

}
func main() {
	o := Order{
		ordId:      456,
		customerId: "39e9e709-dd4f-0512-9488-a67c508b170f",
	}
	reflectInfo(o)
}
