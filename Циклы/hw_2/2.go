package main

import (
	"fmt"
	"time"
)

func main() {
	score := 0
	for i := 0; i < 7; i++ {
		fmt.Println("Едем по уровню")
		fmt.Println("Перепрыгиваем через препятствие")
		fmt.Println("Получено 5 очков")
		score += 5
		fmt.Println("Счет:", score)
		fmt.Println()
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Уровень выполнен, сумма очков за уровень = ", score)
}
