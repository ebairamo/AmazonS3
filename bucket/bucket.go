package bucket

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateContainerName(name string) error {
	if name == "" {
		return fmt.Errorf("имя контейнера не может быть пустым")
	}

	if len(name) < 3 || len(name) > 63 {
		return fmt.Errorf("имя контейнера должно быть от 3 до 63 символов")
	}

	if name[0] == '-' || name[0] == '.' {
		return fmt.Errorf("имя контейнера не может начинаться с дефиса или точки")
	}

	if name[len(name)-1] == '-' || name[len(name)-1] == '.' {
		return fmt.Errorf("имя контейнера не может заканчиваться дефисом или точкой")
	}

	for _, char := range name {
		if !(char >= 'a' && char <= 'z') &&
			!(char >= '0' && char <= '9') &&
			char != '-' && char != '.' {
			return fmt.Errorf("имя контейнера содержит недопустимый символ '%c'", char)
		}

		if char >= 'A' && char <= 'Z' {
			return fmt.Errorf("имя контейнера не может содержать заглавные буквы")
		}
	}

	for i := 0; i < len(name)-1; i++ {
		if (name[i] == '.' && name[i+1] == '.') ||
			(name[i] == '-' && name[i+1] == '-') {
			return fmt.Errorf("имя контейнера содержит последовательные '%c%c'", name[i], name[i+1])
		}
	}

	if isIPAddress(name) {
		return fmt.Errorf("имя контейнера не может быть IP-адресом")
	}

	return nil
}

func isIPAddress(s string) bool {
	parts := strings.Split(s, ".")
	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		for _, char := range part {
			if char < '0' || char > '9' {
				return false
			}
		}

		if len(part) > 1 && part[0] == '0' {
			return false
		}

		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}

	return true
}
