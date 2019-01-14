package faker

import (
	"fmt"
	"testing"
)

func TestAreaCode(t *testing.T) {
	fmt.Printf("Phone Area Code: %s\n", Phone().AreaCode())
}

func TestExchangeCode(t *testing.T) {
	fmt.Printf("Phone Exchange Code: %s\n", Phone().ExchangeCode())
}

func TestSubscriberNumber(t *testing.T) {
	fmt.Printf("Phone Subscriber Number: %s\n", Phone().SubscriberNumber())
}

func TestExtension(t *testing.T) {
	fmt.Printf("Phone Extension: %s\n", Phone().Extension())
}

func TestPhoneFormat(t *testing.T) {
	fmt.Printf("Phone: %s\n", Phone().Format(PhoneFormatExtension))
}

func TestPhoneFormatRandom(t *testing.T) {
	fmt.Printf("Random Phone: %s\n", Phone().Format("{random}"))
}
