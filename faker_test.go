package faker

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	fmt.Printf("Format :: Number: %v; Sequence: %v\n", Format("####"), Format("????"))
}
