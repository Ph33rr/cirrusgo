package runner

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Options struct {
	URL        string
	Proxy      string
	objName    string
	recordId   string
	fwuid      string
	AppName    string
	markup     string
	pageNumber int
	pageSize   int
	Version    bool
	Headers    Headers
	Output     *os.File
}

type Headers []string

func (h Headers) String() string {
	return strings.Join(h, ", ")
}

// Set defines given each header
func (h *Headers) Set(val string) error {
	*h = append(*h, val)
	return nil
}

var o *Options

func init() {

	o = &Options{}

	helpFlagL := flag.String("help", "", "help")
	helpFlagS := flag.String("h", "", "h")
	checkFlagL := flag.String("check", "", "help")
	checkFlagS := flag.String("c", "", "help")
	listFlagL := flag.String("l", "", "help")
	listFlagS := flag.String("list", "", "help")
	urlFlagL := flag.StringVar(&o.URL, "url", "", "")
	urlFlagS := flag.StringVar(&o.URL, "u", "", "")
	headerFlagL := flag.Var(&o.Headers, "header", "")
	headerFlagS := flag.Var(&o.Headers, "header", "")
	proxyFlagL := flag.StringVar(&o.Proxy, "proxy", "", "help")
	versionFlagL := flag.BoolVar(&o.Version, "version", "false", "help")
	versionFlagS := flag.BoolVar(&o.Version, "V", false, "")

	//salesforce app Options
	mainFlagSales := flag.NewFlagSet("salesforce", flag.ExitOnError)
	helpFlagSalesL := mainFlagSales.String("help", "", "help")
	helpFlagSalesS := mainFlagSales.String("h", "", "h")
	checkFlagSalesL := mainFlagSales.String("check", "", "help")
	checkFlagSalesS := mainFlagSales.String("c", "", "help")
	listobjFlagSalesL := mainFlagSales.String("listobj", "", "help")
	listobjFlagSalesS := mainFlagSales.String("lobj", "", "help")
	objFlagSalesL := mainFlagSales.StringVar(&o.objName, "objects", "", "help")
	objFlagSalesS := mainFlagSales.StringVar(&o.objName, "obj", "", "help")
	recordidFlagSalesL := mainFlagSales.StringVar(&o.recordId, "recordid", " ", "help")
	recordidFlagSalesS := mainFlagSales.StringVar(&o.recordId, "re", "", "help")
	fullFlagSalesL := mainFlagSales.String("full", "", "help")
	fullFlagSalesS := mainFlagSales.String("f", "", "help")
	dumpFlagSalesL := mainFlagSales.String("dump", "", "help")
	dumpFlagSalesS := mainFlagSales.String("d", "", "help")
	payloadFlagSalesL := mainFlagSales.String("payload", "", "help")
	pageNumFlagSalesL := mainFlagSales.IntVar(&o.pageNumber, "page", 1, "help")
	pagesFlagSalesL := mainFlagSales.IntVar(&o.pageSize, "pages", 1000, "help")
	fwuidFlagSalesL := mainFlagSales.StringVar(&o.fwuid, "fwuid", "", "help")
	appFlagSalesL := mainFlagSales.StringVar(&o.AppName, "app", "", "help")
	markupFlagSalesL := mainFlagSales.StringVar(&o.markup, "markup", "", "help")

	flag.Usage = func() {
		showBanner()
		h := []string{
			"",
			"Usage:" + usage,
			"",
			"Options:" + options,
			"",
		}

		fmt.Fprint(os.Stderr, strings.Join(h, "\n"))
	}

	flag.Parse()
	mainFlagSales.Parse(os.Args)
}

// ParseOptions will parse given args
func ParseOptions() *Options {
	// Show current version & exit
	if o.Version {
		showVersion()
	}

	// Show banner to user
	// Validate input options
	//o.validate()

	return o
}

