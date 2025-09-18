package spentcalories

import (
	"errors"
	"fmt"
	"log"
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
		log.Println("Ошибка парсинга данных в функции parseTraining(). Неверно переданный формат данных")
		return 0, "", time.Duration(0), errors.New("Ошибка парсинга данных в функции parseTraining()\nНеверно переданный формат данных")
	}
	stepCount, err := strconv.Atoi(splitedData[0])
	if err != nil {
		log.Println(err)
		return 0, "", time.Duration(0), fmt.Errorf("Ошибка парсинга данных в функции parseTraining()\nНе удалось распознать количество шагов\n %w", err)
	}
	if stepCount <= 0 {
		log.Println("Количество шагов меньше либо равно 0")
		return 0, "", time.Duration(0), errors.New("Количество шагов меньше либо равно 0")
	}
	walkingTime, err := time.ParseDuration(splitedData[2])
	if err != nil {
		log.Println(err)
		return 0, "", time.Duration(0), fmt.Errorf("Ошибка парсинга данных в функции parseTraining()\nНе удалось распознать продолжительность\n %w", err)
	}
	if walkingTime <= 0 {
		log.Println("Продолжительность меньше либо равна нулю")
		return 0, "", time.Duration(0), errors.New("Продолжительность меньше либо равна нулю")
	}
	return stepCount, splitedData[1], walkingTime, nil

}

func distance(steps int, height float64) float64 {
	stepLen := height * stepLengthCoefficient
	distance := (float64(steps) * stepLen) / mInKm
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distance := distance(steps, height)
	meanSpeed := distance / duration.Hours()
	return meanSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	stepCount, fitType, fitTime, err := parseTraining(data)
	if err != nil {
		return "", err
	}
	distance := distance(stepCount, height)
	meanSpeed := meanSpeed(stepCount, height, fitTime)

	switch fitType {
	case "Бег":
		calories, err := RunningSpentCalories(stepCount, weight, height, fitTime)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			fitType, fitTime.Hours(), distance, meanSpeed, calories), nil
	case "Ходьба":
		calories, err := WalkingSpentCalories(stepCount, weight, height, fitTime)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			fitType, fitTime.Hours(), distance, meanSpeed, calories), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("В функцию RunningSpentCalories() переданы некорректные данные")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	calories := (duration.Minutes() * weight * meanSpeed) / float64(minInH)
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("В функцию WalkingSpentCalories() переданы некорректные данные")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	calories := (duration.Minutes() * weight * meanSpeed) / float64(minInH)
	return calories * walkingCaloriesCoefficient, nil
}
