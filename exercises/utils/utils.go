package utils

import (
	"fmt"
	"log"
)

func PrintErr(v any) {
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, v)
	log.Println(colored)
}

func PrintInfo(v any) {
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, v)
	log.Println(colored)
}

func PrintSuccess(v any) {
	colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, v)
	log.Println(colored)
}
