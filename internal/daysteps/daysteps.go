package daysteps

import (
	"errors"
	"fmt"
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
		return 0, time.Duration(0), errors.New("Ошибка парсинга данных\nНеверно переданный формат данных")
	}
	stepCount, err := strconv.Atoi(splitedData[0])
	if err != nil {
		return 0, time.Duration(0), fmt.Errorf("Ошибка парсинга данных\nНе удалось распознать количество шагов\n %w", err)
	}
	if stepCount <= 0 {
		return 0, time.Duration(0), errors.New("Количество шагов меньше либо равно 0")
	}
	walkingTime, err := time.ParseDuration(splitedData[1])
	if err != nil {
		return 0, time.Duration(0), fmt.Errorf("Ошибка парсинга данных\nНе удалось распознать продолжительность\n %w", err)
	}
	return stepCount, walkingTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
