package runner


import (
	"fmt"
	"os"
)

func showVersion() {
	fmt.Printf("CirrusGo %s\n", version)
	os.Exit(2)
}
