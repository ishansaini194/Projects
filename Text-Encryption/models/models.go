package models

import (
	"strings"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encrypt(key int, text string) string {
	var builder strings.Builder

	for _, ch := range text {
		pos := strings.IndexRune(alphabet, ch)

		if pos == -1 {
			builder.WriteRune(ch)
			continue
		}

		newPos := (pos + key) % len(alphabet)
		builder.WriteByte(alphabet[newPos])
	}

	return builder.String()
}

func Decrypt(key int, text string) string {
	var builder strings.Builder

	for _, ch := range text {
		pos := strings.IndexRune(alphabet, ch)

		if pos == -1 {
			builder.WriteRune(ch)
			continue
		}

		newPos := (pos - key + len(alphabet)) % len(alphabet)
		builder.WriteByte(alphabet[newPos])
	}

	return builder.String()
}

/*---------- Original Video method ----------

	func hashfun(key int, letter string) (result string) {
		runes := []rune(letter)
		lastLetterKey := string(runes[len(letter)-key : len(letter)])
		leftoverLetters := string(runes[0 : len(letter)-key])
		return fmt.Sprintf("%s%s", lastLetterKey, leftoverLetters)
	}
func Encrypt(key int, plainText string) (result string) {
	hashletter := hashfun(key, originalText)
	var hashedString = ""
	findOne := func(r rune) rune {
		pos := strings.Index(originalText, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalText)) % len(originalText)
			hashedString += string(hashletter[letterPosition])
		}
		return r
	}
	strings.Map(findOne, plainText)
	return hashedString
}
func Decrypt(key int, encryptedText string) (result string) {
	hashLetter := hashfun(key, originalText)
	var hashedString = ""
	findOne := func(r rune) rune {
		pos := strings.Index(hashLetter, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(hashLetter)) % len(hashLetter)
			hashedString += string(originalText[letterPosition])
		}
		return r
	}
	strings.Map(findOne, encryptedText)
	return hashedString
}
*/

/*---------- Alternative implementation using string concatenation ----------
func Encrypt(key int, text string) string {
	result := ""

	for _, ch := range text {
		pos := strings.IndexRune(alphabet, ch)

		if pos == -1 {
			result += string(ch)
			continue
		}

		newPos := (pos + key) % len(alphabet)
		result += string(alphabet[newPos])
	}

	return result
}

func Decrypt(key int, text string) string {
	result := ""

	for _, ch := range text {
		pos := strings.IndexRune(alphabet, ch)

		if pos == -1 {
			result += string(ch)
			continue
		}

		newPos := (pos - key + len(alphabet)) % len(alphabet)
		result += string(alphabet[newPos])
	}

	return result
}
*/
