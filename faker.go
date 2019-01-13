/*Package faker Utility library to generate fake but realistic looking data.
Commonly used for testing, the library can also be used generate canned
data in production.*/
package faker

import (
	"math/rand"
	"time"
)

func init() {
	randomizeSeed()
}

// Format a string replacing # wih a 0-9 number and a ? with a uppercase or lowercase letter.
func Format(str string) string {
	randomizeSeed()
	possibleCharacters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	chars := []rune(str)
	output := make([]rune, 0)
	for _, char := range chars {
		if char == '#' {
			c := '0' + rune(rand.Intn('9'-'0'+1))
			output = append(output, c)
		}

		if char == '?' {
			c := possibleCharacters[rand.Intn(len(possibleCharacters))]
			output = append(output, c)
		}

		if char != '#' && char != '?' {
			output = append(output, char)
		}
	}

	return string(output)
}

func randomizeSeed() {
	rand.Seed(time.Now().UTC().UnixNano())
}
