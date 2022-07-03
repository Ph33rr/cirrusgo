package runner

import (
	"github.com/projectdiscovery/gologger"
)

const banner = `
   ______ _                           ______
  / ____/(_)_____ _____ __  __ _____ / ____/____
 / /    / // ___// ___// / / // ___// / __ / __ \
/ /___ / // /   / /   / /_/ /(__  )/ /_/ // /_/ /
\____//_//_/   /_/    \__,_//____/ \____/ \____/ v0.0.1
`

// Version is the current version of CirrusGo
const Version = `v0.0.1`

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\ttwitter: infosec_90\n\n")
	gologger.Warning().Msgf("Use with caution. You are responsible for your actions\n")
	gologger.Warning().Msgf("Developers assume no liability and are not responsible for any misuse or damage.\n")

}
