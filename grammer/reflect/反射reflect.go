package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int `json:"id" form:"id"`
	Name string
	Age  int
}

func main() {
	u := User{1, "jerry", 23}
	t := reflect.TypeOf(u) //反射出一个interface{}的类型,main.User
	v := reflect.ValueOf(u)
	pv := reflect.ValueOf(&u)
	// 遍历TypeOf 类型
	for i := 0; i < t.NumField(); i++ { //通过索引来取得它的所有字段，这里通过t.NumField来获取它多拥有的字段数量，同时来决定循环的次数
		f := t.Field(i)               //通过这个i作为它的索引，从0开始来取得它的字段
		val := v.Field(i).Interface() //通过interface方法来取出这个字段所对应的值

		//pv.Elem().Field(i).Set(reflect.ValueOf(val))
		fmt.Println(f.Type.String())
		if f.Type.String() == "string" {
			pv.Elem().Field(i).Set(reflect.ValueOf("Name"))
		} else if f.Type.String() == "int" {
			pv.Elem().Field(i).Set(reflect.ValueOf(12))
		}

		fmt.Printf("%6s:%v =%v,tag:%v\n", f.Name, f.Type, val, f.Tag) //Id:int =1,tag:json:"id" form:"id"
	}
}
