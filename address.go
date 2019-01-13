package faker

import (
	"fmt"
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

// CityPrefix generates city name prefix.
func (a *AddressGenerator) CityPrefix() string {
	return cityPrefix[rand.Intn(len(cityPrefix))]
}

// CitySuffix generates city name suffix.
func (a *AddressGenerator) CitySuffix() string {
	return citySuffix[rand.Intn(len(citySuffix))]
}

// City generates a city full name.
func (a *AddressGenerator) City() string {
	option := rand.Intn(8)
	var city string

	switch option {
	case 0:
		city = fmt.Sprintf("%s %s", a.CityPrefix(), Name().First())
	case 1:
		city = fmt.Sprintf("%s %s", a.CityPrefix(), Name().First())
	case 2:
		city = fmt.Sprintf("%s", Name().First())
	case 3:
		city = fmt.Sprintf("%s", Name().Last())
	case 4:
		city = fmt.Sprintf("%s %s%s", a.CityPrefix(), Name().First(), a.CitySuffix())
	case 5:
		city = fmt.Sprintf("%s%s", Name().First(), a.CitySuffix())
	case 6:
		city = fmt.Sprintf("%s%s", Name().Last(), a.CitySuffix())
	case 7:
		city = fmt.Sprintf("%s %s%s", a.CityPrefix(), Name().Last(), a.CitySuffix())
	case 8:
		city = fmt.Sprintf("%s %s", a.CityPrefix(), Name().Last())
	}

	return city
}

// Country generates a country name.
func (a *AddressGenerator) Country() string {
	return countries[rand.Intn(len(countries))]
}

// CountryCode generates a random two letter country code.
func (a *AddressGenerator) CountryCode() string {
	return countryCodes[rand.Intn(len(countryCodes))]
}

// State generates a state name.
func (a *AddressGenerator) State() string {
	return states[rand.Intn(len(states))]
}

// StateCode generates a two letter state code.
func (a *AddressGenerator) StateCode() string {
	return stateCodes[rand.Intn(len(stateCodes))]
}

// StreetSuffix generates the last part of a street name.
func (a *AddressGenerator) StreetSuffix() string {
	return streetSuffix[rand.Intn(len(streetSuffix))]
}

// StreetName generates a random street name.
func (a *AddressGenerator) StreetName() string {
	option := rand.Intn(2)
	var street string

	switch option {
	case 0:
		street = fmt.Sprintf("%s %s", Name().First(), a.StreetSuffix())
	case 1:
		street = fmt.Sprintf("%s %s", Name().Last(), a.StreetSuffix())
	case 2:
		street = fmt.Sprintf("%s %s %s", Name().Last(), Name().First(), a.StreetSuffix())
	}

	return street
}

// Street full street address.
func (a *AddressGenerator) Street() string {
	return fmt.Sprintf("%s %s", a.BuildingNumber(), a.StreetName())
}

// TimeZone generates full time zone name. Might need to be in a different generation category.
func (a *AddressGenerator) TimeZone() string {
	return timeZones[rand.Intn(len(timeZones))]
}

// ZipCode generates standard US zip code.
func (a *AddressGenerator) ZipCode() string {
	return Format("#####")
}

// ZipCodeExtended generates a standard US extended zip code.
func (a *AddressGenerator) ZipCodeExtended() string {
	return Format("#####-####")
}
