package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// создаем структуру с мапой и мьютексом, мбютекс нужен для того, чтобы избежать состояния гонки
// когда к одному и тому же ключу или величине смогут одновременно образаться несколько горутин
type competitiveMap struct {
	mutex sync.Mutex
	m     map[int]int
}

// илициализируем функцию добавления данных. Здесь как раз видим что сначала мьютекс блокирует все каналы кроме того который записывает данные
// потом идет запись и потом разблокировка мьютекса. Важно не забывать его разблокировать, по этому делаем это сразу при помощи дефера
func (cM *competitiveMap) addToMap(key int, value int) {
	cM.mutex.Lock()
	defer cM.mutex.Unlock()

	cM.m[key] = value
}

func main() {
	//инициализируем переменную
	storage := competitiveMap{
		mutex: sync.Mutex{},
		m:     map[int]int{},
	}

	//создаем вейт группу
	wg := sync.WaitGroup{}
	wg.Add(5)

	//записываем в мапу рандомные данные с рандомным ключем
	for i := 0; i < 5; i++ {
		go func(wg *sync.WaitGroup) {
			storage.addToMap(rand.Intn(100), rand.Intn(100))
			wg.Done()
		}(&wg)
	}

	wg.Wait()

	//посмотрим что у нас получилось
	fmt.Println(storage.m)
}
