package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Игра 'Угадай число' - от 1 до 100 началась!")
	fmt.Println("Угадайте число за 10 попыток!")
	var history []int
	var attempts int = 10
	var guess int // число которое  ввел пользователь
	var diff int  // разница между тем что ввел пользователь и загаданным числом
	secret := randomInt()
	for {
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Ошибка: некорректное число")
			fmt.Scanln() // чистит буфер от букв
			continue     // бросает выполнение круга и возвращается наверх к следующей попытке
			history = append(history, guess)
		}
		if guess > secret {
			diff = guess - secret
		} else {
			diff = secret - guess
		}
		if guess == secret {
			fmt.Println("Поздравляю, ты угадал число")
			break
		} else {
			if diff <= 5 {
				fmt.Println("🔥 Горячо")
				attempts--
				history = append(history, guess)
				fmt.Println("Ранее введённые числа:", history)
			} else if diff <= 15 {
				fmt.Println("🙂 Тепло")
				attempts--
				history = append(history, guess)
				fmt.Println("Ранее введённые числа:", history)
			} else {
				fmt.Println("❄️ Холодно")
				fmt.Println("Неверное число")
				history = append(history, guess)
				fmt.Println("Ранее введённые числа:", history)
				attempts--

			}
			if attempts == 0 {
				fmt.Println(secret, "это было загаданное число")
				return
			}
		}
	}
}

func randomInt() int {
	n := rand.Intn(101)
	return n
}
