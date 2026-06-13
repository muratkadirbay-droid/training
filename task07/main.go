package main

import (
	"fmt"
	"math"
)

func main() {
	bannerWidth := 12
	bannerHeight := 8

	bannerArea := bannerWidth * bannerHeight
	fmt.Println(bannerArea)

	halfBannerArea := bannerArea / 2
	fmt.Println(halfBannerArea)

	bannerBorderLength := (bannerWidth + bannerHeight) * 2
	fmt.Println(bannerBorderLength)

	boxCount := 29
	leftoverBoxes := boxCount % 5
	fmt.Println(leftoverBoxes)

	tempMorning := 18
	tempAfternoon := 27
	tempEvening := 21

	totalTemp := tempMorning + tempAfternoon + tempEvening
	averageTemp := totalTemp / 3

	fmt.Println(averageTemp)

	knownWords := 47
	wordsGoal := 120

	progressPercent := float64(knownWords) / float64(wordsGoal) * 100
	fmt.Println(progressPercent, "%")

	coins := 0

	coins += 500
	fmt.Println(coins)

	coins += 1200
	fmt.Println(coins)

	coins /= 2
	fmt.Println(coins)

	coins *= 2
	fmt.Println(coins)

	coins -= 300
	fmt.Println(coins)

	participants := 42
	groupCount := 8

	participantsPerGroup := participants / groupCount
	fmt.Println(participantsPerGroup)

	fmt.Println(20 - 4*3)
	fmt.Println((20 - 4) * 3)


	squareValue := 81.0
	fmt.Println(math.Sqrt(squareValue))

	multiplier := 5.0
	exponent := 2.0

	fmt.Println(math.Pow(multiplier, exponent))
}