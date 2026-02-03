package utils

import (
	"fmt"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
