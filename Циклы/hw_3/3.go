package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	score := 0
	for {
		fmt.Println("Едем по уровню")
		if rand.Intn(8) == 1 {
			fmt.Println()
			fmt.Println("Врезался в препятствие(")
			break
		}
		fmt.Println("Перепрыгиваем через препятствие")
		fmt.Println("Получено 5 очков")
		score += 5
		fmt.Println("Счет:", score)
		fmt.Println()
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Уровень завершен, сумма очков за уровень = ", score)
}
