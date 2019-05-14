package main

import (
	"bytes"
	"fmt"
)

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	// WriteByte func(c byte) error
	// WriteByte appends the byte c to the buffer, growing the buffer as needed. The returned error is always nil, but is included to match bufio.Writer's WriteByte. If the buffer becomes too large, WriteByte will panic with ErrTooLarge.
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}
