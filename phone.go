package faker

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	// PhoneFormatFull ({area}) {exchange}-{subscriber}
	PhoneFormatFull = "({area}) {exchange}-{subscriber}"
	// PhoneFormatExtension ({area}) {exchange}-{wubscriber} #{extension}
	PhoneFormatExtension = "({area}) {exchange}-{subscriber} #{extension}"
	// PhoneFormat {area}{exchange}{subscriber}
	PhoneFormat = "{area}{exchange}{subscriber}"
	// PhoneFormatRandom generates a random name combination.
	PhoneFormatRandom = "{random}"
)

// PhoneGenerator container for phone generation functions.
type PhoneGenerator struct{}

var phone = PhoneGenerator{}

// Phone generates parts of a US phone number.
func Phone() *PhoneGenerator {
	return &phone
}

/*
  This follows the rules outlined in the North American Numbering Plan
  at https://en.wikipedia.org/wiki/North_American_Numbering_Plan.
  The NANP number format may be summarized in the notation NPA-NXX-xxxx:

*/

/*AreaCode The allowed ranges for NPA (area code) are: [2–9] for the first digit, and
[0-9] for the second and third digits. The NANP is not assigning area codes
with 9 as the second digit.*/
func (p *PhoneGenerator) AreaCode() string {
	return fmt.Sprintf("%d%d%d", digit(2, 9), digit(0, 8), digit(0, 9))
}

/*ExchangeCode The allowed ranges for NXX (central office/exchange) are: [2–9] for the first
digit, and [0–9] for both the second and third digits (however, in geographic
area codes the third digit of the exchange cannot be 1 if the second digit is
also 1).*/
func (p *PhoneGenerator) ExchangeCode() string {
	secondDigit := digit(0, 9)
	var thirdDigit int8

	switch secondDigit {
	case 1:
		thirdDigit = digit(2, 9)
	default:
		thirdDigit = digit(0, 9)
	}

	return fmt.Sprintf("%d%d%d", digit(2, 9), secondDigit, thirdDigit)
}

// SubscriberNumber The allowed ranges for xxxx (subscriber number) are [0–9] for each of the four digits.
func (p *PhoneGenerator) SubscriberNumber() string {
	return Format("####")
}

// Extension genrates a random sized extension number.
func (p *PhoneGenerator) Extension() string {
	options := []string{"#####", "####", "###", "##", "#"}
	selected := options[rand.Intn(len(options))]
	return Format(selected)
}

/*Format takes in a format and generates the phone components asked for.
Formatters:

{area}, {exchange}, {subscriber}, {extension}

{random} - uses a random prefabbed phone combination
*/
func (p *PhoneGenerator) Format(format string) string {
	val := format

	if strings.Contains(val, "{area}") {
		val = strings.Replace(val, "{area}", p.AreaCode(), -1)
	}

	if strings.Contains(val, "{exchange}") {
		val = strings.Replace(val, "{exchange}", p.ExchangeCode(), -1)
	}

	if strings.Contains(val, "{subscriber}") {
		val = strings.Replace(val, "{subscriber}", p.SubscriberNumber(), -1)
	}

	if strings.Contains(val, "{extension}") {
		val = strings.Replace(val, "{extension}", p.Extension(), -1)
	}

	if strings.Contains(val, "{random}") {
		possible := []string{PhoneFormatFull, PhoneFormatExtension, PhoneFormat}
		val = strings.Replace(val, "{random}", p.Format(possible[rand.Intn(len(possible))]), -1)
	}

	return val
}
