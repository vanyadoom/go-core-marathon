package main

import "fmt"

func Generator() <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func Validator(in <-chan int) <-chan int { // 🟢 ИСПРАВЛЕНО: Задали возвращаемый тип
	out := make(chan int)
	go func() { // 🟢 ИСПРАВЛЕНО: Упаковали в фоновую горутину
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
		close(out) // 🟢 ИСПРАВЛЕНО: Каскадно закрываем следующую трубу
	}()
	return out
}

func Calculator(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 100
		}
		close(out)
	}()
	return out
}

func main() {
	stage1 := Generator()
	stage2 := Validator(stage1)
	stage3 := Calculator(stage2)

	for result := range stage3 {
		fmt.Println("Конвейер выдал чистый платёж:", result)
	}
}
