package runner

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Headers defines custom headers
type Headers []string

func (h Headers) String() string {
	return strings.Join(h, ", ")
}

// Set defines given each header
func (h *Headers) Set(val string) error {
	*h = append(*h, val)
	return nil
}
