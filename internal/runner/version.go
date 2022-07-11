package runner

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

func showVersion() {
	fmt.Printf("\n[%s] CirrusGo %s\n", aurora.Green("VER").String(),
		aurora.White(version).String())
	os.Exit(2)
}
