package main

import "fmt"

//Адаптер - как понятно из названия адаптирует, в нашем случае он позволяет использовать несовместимый интерфейс в нужном месте

// интерфейс который будет адаптировать, сам адаптер по сути
type Adapter interface {
	Operation()
}

// а это адаптируеМЫЙ, его метод нужно будет где-то вызывать
type AdaptableClass struct {
}

// сам его метод
func (adaptee *AdaptableClass) AdaptedOperation() {
	fmt.Println("Это то, что мы хотим адаптировать под что-то свое")
}

// обертка для адаптируемого класса
type ConcreteAdapter struct {
	*AdaptableClass
}

// метод который будет удовлетворять тому, что мы хотим использовать. То есть тут мы преобразуем метод который получаем и НЕ можем использовать
// в тот который сможем
func (adapter *ConcreteAdapter) Operation() {
	adapter.AdaptedOperation()
}

// в конструкторе уже преобразуем и адаптируем интерфейс
func NewAdapter(adaptee *AdaptableClass) Adapter {
	return &ConcreteAdapter{adaptee}
}

// пробуем применить
func main() {
	adapter := NewAdapter(&AdaptableClass{})
	adapter.Operation()
}
