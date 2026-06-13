package main

import "fmt"

func main() {
	cpuLoad := map[int]int{
		1: 40,
		2: 65,
		3: 30,
	}

	maxCore := 0
	maxLoad := -1

	for core, load := range cpuLoad {
		if load > maxLoad {
			maxLoad = load
			maxCore = core
		}
	}

	fmt.Println("Максимальная загрузка у ядра:", maxCore)
}