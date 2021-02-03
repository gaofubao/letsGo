package stringexp

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrings(t *testing.T)  {
	var b strings.Builder
	fmt.Println(b)

	var r strings.Reader
	fmt.Println(r)

	var re strings.Replacer
	fmt.Println(re)
}
