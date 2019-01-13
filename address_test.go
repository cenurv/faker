package faker

import (
	"fmt"
	"testing"
)

func TestBuildingNumber(t *testing.T) {
	fmt.Printf("Building number: %v\n", Address().BuildingNumber())
}

func TestCityPrefix(t *testing.T) {
	fmt.Printf("City Prefix: %s\n", Address().CityPrefix())
}

func TestCitySuffix(t *testing.T) {
	fmt.Printf("City Suffix: %s\n", Address().CitySuffix())
}

func TestCity(t *testing.T) {
	fmt.Printf("City: %s\n", Address().City())
}

func TestCountry(t *testing.T) {
	fmt.Printf("Country: %s\n", Address().Country())
}

func TestCountryCode(t *testing.T) {
	fmt.Printf("CountryCode: %s\n", Address().CountryCode())
}

func TestState(t *testing.T) {
	fmt.Printf("State: %s\n", Address().State())
}

func TestStateCode(t *testing.T) {
	fmt.Printf("StateCode: %s\n", Address().StateCode())
}

func TestStreetSuffix(t *testing.T) {
	fmt.Printf("StreetSuffix: %s\n", Address().StreetSuffix())
}

func TestStreetName(t *testing.T) {
	fmt.Printf("StreetName: %s\n", Address().StreetName())
}

func TestStreet(t *testing.T) {
	fmt.Printf("Street: %s\n", Address().Street())
}

func TestTimeZone(t *testing.T) {
	fmt.Printf("TimeZone: %s\n", Address().TimeZone())
}

func TestZipCode(t *testing.T) {
	fmt.Printf("ZipCode: %s\n", Address().ZipCode())
}

func TestZipCodeExtended(t *testing.T) {
	fmt.Printf("ZipCodeExtended: %s\n", Address().ZipCodeExtended())
}
