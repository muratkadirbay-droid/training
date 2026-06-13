package main

import "fmt"

func main() {
	runningExercises := []string{"Sprint", "Interval Running"}
	walkingExercises := []string{"Power Walking", "Nordic Walking"}

	fmt.Println("Бег:", len(runningExercises))
	fmt.Println("Ходьба:", len(walkingExercises))
}