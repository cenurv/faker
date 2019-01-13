package faker

import (
	"fmt"
	"testing"
)

func TestFirstName(t *testing.T) {
	fmt.Printf("First Name: %s\n", Name().First())
}

func TestMiddleName(t *testing.T) {
	fmt.Printf("Middle Name: %s\n", Name().Middle())
}

func TestLasttName(t *testing.T) {
	fmt.Printf("Last Name: %s\n", Name().Last())
}

func TestInitial(t *testing.T) {
	fmt.Printf("Inital: %s\n", Name().Initial())
}

func TestPrefix(t *testing.T) {
	fmt.Printf("Prefix: %s\n", Name().Prefix())
}

func TestSuffix(t *testing.T) {
	fmt.Printf("Suffix: %s\n", Name().Suffix())
}

func TestFormattedName(t *testing.T) {
	fmt.Printf("Name Format: %s\n", Name().Format("{prefix} {first} {middle} {initial} {last} {suffix}, {title}"))
}

func TestFormattedRandomName(t *testing.T) {
	fmt.Printf("Random Name Format: %s\n", Name().Format(NameFormatRandom))
}

func TestStandardName(t *testing.T) {
	fmt.Printf("Standard Name: %s\n", Name().Standard())
}

func TestFullName(t *testing.T) {
	fmt.Printf("Full Name: %s\n", Name().Full())
}
