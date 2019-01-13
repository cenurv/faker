package faker

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	// NameFormatFull {first} {middle} {last}
	NameFormatFull = "{first} {middle} {last}"
	// NameFormatMiddleInitial {first} {initial} {last}
	NameFormatMiddleInitial = "{first} {initial} {last}"
	// NameFormat {first} {last}
	NameFormat = "{first} {last}"
	// NameFormatRandom generates a random name combination.
	NameFormatRandom = "{random}"
)

// NameGenerator container for name generation functions.
type NameGenerator struct{}

var n = NameGenerator{}

// Name generates parts of a person's name.
func Name() *NameGenerator {
	return &n
}

// First generates a first name.
func (n *NameGenerator) First() string {
	return firstNames[rand.Intn(len(firstNames))]
}

// Middle generates a first name.
func (n *NameGenerator) Middle() string {
	return firstNames[rand.Intn(len(firstNames))]
}

// Initial generates a single letter initial.
func (n *NameGenerator) Initial() string {
	return fmt.Sprintf("%s.", strings.ToUpper(Format("?")))
}

// Last generates a last name.
func (n *NameGenerator) Last() string {
	return lastNames[rand.Intn(len(lastNames))]
}

// Prefix generates a name prefix.
func (n *NameGenerator) Prefix() string {
	return prefixs[rand.Intn(len(prefixs))]
}

// Suffix generates a name prefix.
func (n *NameGenerator) Suffix() string {
	return suffixs[rand.Intn(len(suffixs))]
}

/*Format takes in a format and generates the name components asked for.
Formatters:

{first}, {middle}, {last}, {initial}, {prefix}, {suffix}, {title}

{random} - uses a random prefabbed name combination
*/
func (n *NameGenerator) Format(format string) string {
	val := format

	if strings.Contains(val, "{prefix}") {
		val = strings.Replace(val, "{prefix}", n.Prefix(), -1)
	}

	if strings.Contains(val, "{first}") {
		val = strings.Replace(val, "{first}", n.First(), -1)
	}

	if strings.Contains(val, "{middle") {
		val = strings.Replace(val, "{middle}", n.First(), -1)
	}

	if strings.Contains(val, "{initial") {
		val = strings.Replace(val, "{initial}", n.Initial(), -1)
	}

	if strings.Contains(val, "{last}") {
		val = strings.Replace(val, "{last}", n.Last(), -1)
	}

	if strings.Contains(val, "{suffix}") {
		val = strings.Replace(val, "{suffix}", n.Suffix(), -1)
	}

	if strings.Contains(val, "{title}") {
		val = strings.Replace(val, "{title}", n.Title(), -1)
	}

	if strings.Contains(val, "{random}") {
		possible := []string{NameFormatFull, NameFormatMiddleInitial, NameFormat}
		val = strings.Replace(val, "{random}", n.Format(possible[rand.Intn(len(possible))]), -1)
	}

	return val
}

// Standard {first} {last} name format.
func (n *NameGenerator) Standard() string {
	return n.Format(NameFormat)
}

// Full {first} {last} name format.
func (n *NameGenerator) Full() string {
	return n.Format(NameFormatFull)
}

// TitleDescriptor generates a job title descriptor.
func (n *NameGenerator) TitleDescriptor() string {
	return titleDescriptors[rand.Intn(len(titleDescriptors))]
}

// TitleJob generates a job title job.
func (n *NameGenerator) TitleJob() string {
	return titleJobs[rand.Intn(len(titleJobs))]
}

// TitleLevel generates a job title descriptor.
func (n *NameGenerator) TitleLevel() string {
	return titleLevels[rand.Intn(len(titleLevels))]
}

// Title generates a full job title.
func (n *NameGenerator) Title() string {
	return fmt.Sprintf("%s %s %s", n.TitleDescriptor(), n.TitleLevel(), n.TitleJob())
}
