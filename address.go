package faker

import (
	"math/rand"
)

// Address generator for common address information.
type address struct{}

func Address() *address {
	return &address{}
}

func (a *address) BuildingNumber() string {
	randomizeSeed()
	options := []string{"#####", "####", "###", "##", "#"}
	selected := options[rand.Intn(len(options))]
	return Format(selected)
}
