package daysteps

import (
	"errors"
	"fmt"
	"spentcalories"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	splitedData := strings.Split(data, ",")
	if len(splitedData) != 2 {
		return 0, time.Duration(0), errors.New("Ошибка парсинга данных в функции parsePackage()\nНеверно переданный формат данных")
	}
	stepCount, err := strconv.Atoi(splitedData[0])
	if err != nil {
		return 0, time.Duration(0), fmt.Errorf("Ошибка парсинга данных в функции parsePackage()\nНе удалось распознать количество шагов\n %w", err)
	}
	if stepCount <= 0 {
		return 0, time.Duration(0), errors.New("Количество шагов меньше либо равно 0")
	}
	walkingTime, err := time.ParseDuration(splitedData[1])
	if err != nil {
		return 0, time.Duration(0), fmt.Errorf("Ошибка парсинга данных в функции parsePackage()\nНе удалось распознать продолжительность\n %w", err)
	}
	return stepCount, walkingTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	stepCount, walkingTime, err := parsePackage(data)
	if err != nil {
		fmt.Print(err)
		return ""
	}
	if err == errors.New("Количество шагов меньше либо равно 0") {
		return ""
	}
	distance := (float64(stepCount) * stepLength) / mInKm
	calories, err := spentcalories.WalkingSpentCalories(stepCount, weight, height, walkingTime)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f.\nВы сожгли %.2f ккал.", stepCount, distance, calories)
}
