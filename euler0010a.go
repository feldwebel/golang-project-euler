package main

import (
	"fmt"
	"runtime"
	"sync"

	"golang-project-euler/common"
)

const limit = 2_000_000

func main() {
	numWorkers := runtime.NumCPU()

	// Канал для сбора простых чисел от воркеров
	results := make(chan uint64, 1024)

	// Канал с заданиями (нечётные числа для проверки)
	jobs := make(chan uint64, 1024)

	var wg sync.WaitGroup

	// Запускаем воркеров
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range jobs {
				if common.IsPrime(n) {
					results <- n
				}
			}
		}()
	}

	// Генератор заданий в отдельной горутине
	go func() {
		for i := uint64(3); i < limit; i += 2 {
			jobs <- i
		}
		close(jobs)
	}()

	// Закрываем канал результатов, когда все воркеры закончат
	go func() {
		wg.Wait()
		close(results)
	}()

	// Суммируем результаты (начинаем с 2 — единственного чётного простого)
	var sum uint64 = 2
	for p := range results {
		sum += p
	}

	fmt.Println(sum)
}
