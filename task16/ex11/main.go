package main

import "fmt"

func main() {
	defaultConfig := map[string]string{
		"host": "localhost",
		"port": "8080",
		"mode": "production",
	}

	currentConfig := map[string]string{
		"host": "localhost",
		"port": "8080",
		"mode": "production",
	}

	same := true

	for k, v := range defaultConfig {
		if currentConfig[k] != v {
			same = false
			break
		}
	}

	if same {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}

	// изменение mode
	currentConfig["mode"] = "debug"

	same = true
	for k, v := range defaultConfig {
		if currentConfig[k] != v {
			same = false
			break
		}
	}

	if same {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}
}