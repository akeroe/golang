package main

import (
	"fmt"
	"math/rand"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func main() {
	for {

		fmt.Println("Игра 'Угадай число' - выберите режим сложности: 1- легкий, 2-средний, 3-сложный")

		var history []int // тут слайс введенных ранее чисел пользователем
		var attempts int = 10
		var guess int // число, которое ввел пользователь
		var diff int  // разница между тем, что ввел пользователь, и загаданным числом
		var maxNumberInt int
		choice := checkchoice() // выбор сложности пользователя

		maxNumberInt, attempts = choiceLevelOfDifficult(choice)
		secret := randomInt(maxNumberInt, choice)

		for {
			_, err := fmt.Scan(&guess)
			if err != nil {
				fmt.Println("Ошибка: некорректное число")
				fmt.Scanln() // чистит буфер от букв
				continue     // бросает выполнение круга и возвращается наверх
			}

			// Считаем разницу
			if guess > secret {
				diff = guess - secret
			} else {
				diff = secret - guess
			}

			// Проверяем результат
			if guess == secret {
				fmt.Println(Green + "Поздравляю, ты угадал число" + Reset)
				break
			} else {
				//подсказки
				if diff <= 5 {
					fmt.Println("🔥 Горячо")
				} else if diff <= 15 {
					fmt.Println("🙂 Тепло")
				} else {
					fmt.Println("❄️ Холодно")
				}

				// 2. общ код выполняется для ЛЮБОГО неверного ответа
				attempts--
				history = append(history, guess)
				fmt.Println(Yellow+"Ранее введённые числа:"+Reset, history)

				// 3. Проверяем, не закончились ли попытки
				if attempts == 0 {
					fmt.Println(Red+"Проигрыш\n", secret, "— это было загаданное число"+Reset)
					break
				}
			}
		}
		fmt.Scanln()
		fmt.Println("Хотите сыграть еще раз? 1 - ДА, 0 - НЕТ")
		var playAgain int
		fmt.Scanln(&playAgain)
		if playAgain != 1 {
			fmt.Println("Спасибо за игру. Пока!")
			break
		}
	}
}

func randomInt(maxNumberInt, choice int) int { //генерирует рандом число
	return rand.Intn(maxNumberInt)
}
func choiceLevelOfDifficult(choice int) (maxNumberInt int, attempts int) {

	switch choice {
	case 1:
		maxNumberInt = 50
		attempts = 15
		fmt.Println("Выбран легкий уровень: числа 0 - 50, 15 попыток\nИгра началась: вводи число")

	case 2:
		maxNumberInt = 100
		attempts = 10
		fmt.Println("Выбран средний уровень: 0 - 100, 10 попыток\nИгра началась: вводи число")
	case 3:
		maxNumberInt = 200
		attempts = 5
		fmt.Println("Выбран сложный уровень: 0 - 200, 5 попыток\nИгра началась: вводи число")
	default:
		fmt.Println("Ошибка: Введите 1, 2 или 3")

	}
	return maxNumberInt, attempts
}
func checkchoice() int {
	var choice int
	for {
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Ошибка: введите число!")
			fmt.Scanln()
			continue
		}
		if choice < 1 || choice > 3 {
			fmt.Println("Ошибка: Введите от 1 до 3 ")
			continue
		}
		return choice
	}
}
