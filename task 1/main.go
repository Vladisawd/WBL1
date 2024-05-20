package main

type Human struct {
	name string
	age  int
}

type Action struct {
	Human
	run  string
	stop string
}

func main() {
	a := Action{}
	a.name = "Aboba" //Видно что не смотря на отсутствие конкретных переменных name, age в структуре Action, мы можем их использовать из за аггрегации
	a.age = 54
}
