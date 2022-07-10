package runner


import (
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

func showBanner() {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\ttwitter: infosec_90\n\n")
	gologger.Warning().Msg("Use with caution. You are responsible for your actions\n")
	gologger.Warning().Msg("Developers assume no liability and are not responsible for any misuse or damage.\n")
}
