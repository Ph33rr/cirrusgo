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

	flag.String("help", "flag.ExitOnError", "")
	flag.String("h", "", "h")
	flag.String("check", "", "help")
	flag.String("c", "", "help")
	flag.String("l", "", "help")
	flag.String("list", "", "help")
	flag.StringVar(&o.URL, "url", "", "")
	flag.StringVar(&o.URL, "u", "", "")
	flag.Var(&o.Headers, "header", "")
	flag.Var(&o.Headers, "header", "")
	flag.StringVar(&o.Proxy, "proxy", "", "help")
	flag.BoolVar(&o.Version, "version", false, "help")
	flag.BoolVar(&o.Version, "V", false, "")

	//salesforce app Options
	mainFlagSales := flag.NewFlagSet("salesforce", flag.ExitOnError)
	mainFlagSales.String("help", "", "help")
	mainFlagSales.String("h", "", "h")
	mainFlagSales.String("check", "", "help")
	mainFlagSales.String("c", "", "help")
	mainFlagSales.String("listobj", "", "help")
	mainFlagSales.String("lobj", "", "help")
	mainFlagSales.StringVar(&o.objName, "objects", "", "help")
	mainFlagSales.StringVar(&o.objName, "obj", "", "help")
	mainFlagSales.StringVar(&o.recordId, "recordid", " ", "help")
	mainFlagSales.StringVar(&o.recordId, "re", "", "help")
	mainFlagSales.String("full", "", "help")
	mainFlagSales.String("f", "", "help")
	mainFlagSales.String("dump", "", "help")
	mainFlagSales.String("d", "", "help")
	mainFlagSales.String("payload", "", "help")
	mainFlagSales.IntVar(&o.pageNumber, "page", 1, "help")
	mainFlagSales.IntVar(&o.pageSize, "pages", 1000, "help")
	mainFlagSales.StringVar(&o.fwuid, "fwuid", "", "help")
	mainFlagSales.StringVar(&o.AppName, "app", "", "help")
	mainFlagSales.StringVar(&o.markup, "markup", "", "help")
	
	
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

