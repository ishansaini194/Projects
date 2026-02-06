package main

import (
	"fmt"

	"github.com/ishansaini194/Projects/models"
)

func main() {
	plainText := "HELLOWORLD"
	fmt.Println("Plain text:", plainText)

	encryption := models.Encrypt(5, plainText)
	fmt.Println("Encrypted text:", encryption)

	decrypted := models.Decrypt(5, encryption)
	fmt.Println("Decrypted:", decrypted)
}
