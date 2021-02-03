package bytesexp

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBytes(t *testing.T)  {
	var b bytes.Buffer
	fmt.Println(b)

	var r bytes.Reader
	fmt.Println(r)
}
