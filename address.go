package faker

import (
	"math/rand"
)

// AddressGenerator is a container for address generator functions.
type AddressGenerator struct{}

var a = AddressGenerator{}

// Address generator for common address information.
func Address() *AddressGenerator {
	return &a
}

// BuildingNumber generates a random building number.
func (a *AddressGenerator) BuildingNumber() string {
	options := []string{"#####", "####", "###", "##", "#"}
	selected := options[rand.Intn(len(options))]
	return Format(selected)
}
