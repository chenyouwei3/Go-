package main

import "fmt"

type ZhiFuBao struct{}

type WeChat struct{}

type Payer interface {
	Pay(int64)
}

func (w *WeChat) Pay(amount int64) {
	fmt.Printf("使用微信付款：%.2f元。\n", float64(amount/100))
}

func (z *ZhiFuBao) Pay(amount int64) {
	fmt.Printf("使用支付宝付款：%.2f元。\n", float64(amount/100))
}

func Checkout(obj Payer) {
	obj.Pay(100)
}

//func CheckoutWithZFB(obj *ZhiFuBao) {
//	obj.Pay(100)
//}
//
//func CheckoutWithWX(obj *WeChat) {
//	obj.Pay(100)
//}

func main() {
	Checkout(&ZhiFuBao{})
}
