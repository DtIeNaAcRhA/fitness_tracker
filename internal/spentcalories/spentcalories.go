package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	splitedData := strings.Split(data, ",")
	if len(splitedData) != 3 {
		return 0, "", time.Duration(0), errors.New("Ошибка парсинга данных в функции parseTraining()\nНеверно переданный формат данных")
	}
	stepCount, err := strconv.Atoi(splitedData[0])
	if err != nil {
		return 0, "", time.Duration(0), fmt.Errorf("Ошибка парсинга данных в функции parseTraining()\nНе удалось распознать количество шагов\n %w", err)
	}
	if stepCount <= 0 {
		return 0, "", time.Duration(0), errors.New("Количество шагов меньше либо равно 0")
	}
	walkingTime, err := time.ParseDuration(splitedData[2])
	if err != nil {
		return 0, "", time.Duration(0), fmt.Errorf("Ошибка парсинга данных в функции parseTraining()\nНе удалось распознать продолжительность\n %w", err)
	}
	return stepCount, splitedData[1], walkingTime, nil

}

func distance(steps int, height float64) float64 {
	stepLen := height * stepLengthCoefficient
	distance := (float64(steps) * stepLen) / mInKm
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}
