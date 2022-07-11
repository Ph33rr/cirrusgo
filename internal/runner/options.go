package runner

import (
	"CirrusGo/pkg/errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Options struct {
	URL              string
	List             string
	Proxy            string
	objName          string
	recordId         string
	fwuid            string
	AppName          string
	markup           string
	Target           string
	infosec_90       string
	listApp          bool
	check            bool
	listobj          bool
	getobj           bool
	getrecord        bool
	checkWritableOBJ bool
	full             bool
	dump             bool
	help             bool
	payload          bool
	GetItems         bool
	GetRecord        bool
	WritableOBJ      bool
	SearchObj        bool
	AuraContext      bool
	ObjectList       bool
	Dump             bool
	pageNumber       int
	pageSize         int
	Version          bool
	Headers          Headers
	Output           *os.File
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

	flag.Usage = func() {
		showBanner()
		h := []string{
			"",
			"Usage:" + usage,
			"",
			"Options:" + optionss,
			"",
		}
		fmt.Fprint(os.Stderr, strings.Join(h, "\n"))
	}

}

func ParseOptions() *Options {
	// Show current version & exit
	o = &Options{}

	flag.BoolVar(&o.help, "help", false, "")
	flag.BoolVar(&o.help, "h", false, "")
	flag.String("check", "", "help")
	flag.String("c", "", "help")
	flag.String("l", "", "help")
	flag.BoolVar(&o.listApp, "list", false, "Show App Support")
	flag.StringVar(&o.URL, "url", "", "Define single URL")
	flag.StringVar(&o.URL, "u", "", "Define single URL")
	flag.StringVar(&o.infosec_90, "o", "", "")
	flag.StringVar(&o.infosec_90, "output", "", "")
	flag.Var(&o.Headers, "header", "")
	flag.Var(&o.Headers, "H", "")
	flag.StringVar(&o.Proxy, "proxy", "", "help")
	flag.BoolVar(&o.Version, "version", false, "help")
	flag.BoolVar(&o.Version, "V", false, "")

	//salesforce app Options
	mainFlagSales := flag.NewFlagSet("salesforce", flag.ContinueOnError)
	mainFlagSales.BoolVar(&o.help, "help", false, "Display its help")
	mainFlagSales.BoolVar(&o.help, "h", false, "Display its help")
	mainFlagSales.StringVar(&o.URL, "u", "", "Define single URL")
	mainFlagSales.StringVar(&o.URL, "url", "", "Define single URL")
	mainFlagSales.BoolVar(&o.check, "check", false, "only check endpoint")
	mainFlagSales.BoolVar(&o.check, "c", false, "only check endpoint")
	mainFlagSales.BoolVar(&o.listobj, "listobj", false, " pull the object list")
	mainFlagSales.BoolVar(&o.listobj, "lobj", false, " pull the object list")
	mainFlagSales.BoolVar(&o.getobj, "getobj", false, "pull the object")
	mainFlagSales.BoolVar(&o.getobj, "gobj", false, " pull the object")
	mainFlagSales.StringVar(&o.objName, "objects", "User", "set the object name. Default value is 'User' objectn")
	mainFlagSales.StringVar(&o.objName, "obj", "User", "set the object name. Default value is User object")
	mainFlagSales.BoolVar(&o.getrecord, "getrecord", false, "set the recode id to dump the record")
	mainFlagSales.BoolVar(&o.getrecord, "gre", false, "set the recode id to dump the record")
	mainFlagSales.StringVar(&o.recordId, "recordid", "RecordID", "set the recode id to dump the record")
	mainFlagSales.StringVar(&o.recordId, "re", "RecordID", "set the recode id to dump the record")
	mainFlagSales.BoolVar(&o.checkWritableOBJ, "chkWritable ", false, " check all Writable objects")
	mainFlagSales.BoolVar(&o.checkWritableOBJ, "cw", false, " check all Writable objects")
	mainFlagSales.BoolVar(&o.full, "full", false, "dump all pages of objects.")
	mainFlagSales.BoolVar(&o.full, "f", false, " dump all pages of objects.")
	mainFlagSales.BoolVar(&o.dump, "dump", false, "get json file")
	mainFlagSales.BoolVar(&o.dump, "d", false, "get json file")
	mainFlagSales.BoolVar(&o.payload, "payload", false, "(case-sensitive) Generator payload for test manual Default 'ObjectList'")
	mainFlagSales.BoolVar(&o.GetItems, "GetItems", false, "Generator payload for test manual Default 'User'")
	mainFlagSales.BoolVar(&o.ObjectList, "ObjectList", false, "Generator payload for test manual Default 'ObjectList'")
	mainFlagSales.BoolVar(&o.GetRecord, "GetRecord", false, "Generator payload for test manual Default 'Ph33rr'")
	mainFlagSales.BoolVar(&o.WritableOBJ, "WritableOBJ", false, "Generator payload for test manual Default 'User'")
	mainFlagSales.BoolVar(&o.SearchObj, "SearchObj", false, "Generator payload for test manual Default 'User'")
	mainFlagSales.BoolVar(&o.AuraContext, "AuraContext", false, "Generator payload for test manual")
	mainFlagSales.BoolVar(&o.Dump, "Dump", false, "Generator payload for test manual")
	mainFlagSales.IntVar(&o.pageNumber, "page", 1, "page set page")
	mainFlagSales.IntVar(&o.pageSize, "pages", 1000, "pages set pageSize")
	mainFlagSales.StringVar(&o.fwuid, "fwuid", "fwuidHere", "fwuid set UID ")
	mainFlagSales.StringVar(&o.AppName, "app", "AppNamehere", "App set AppName")
	mainFlagSales.StringVar(&o.markup, "markup", "markuphere", "markup set markup")

	flag.Parse()
	if len(os.Args) != 1 {
		if AppNamevalidate(os.Args[1]) {
			mainFlagSales.Parse(os.Args[2:])
		}
	}
	o.validate()
	return o
}

func (o *Options) validate() {

	if isStdin() {
		b, e := ioutil.ReadAll(os.Stdin)
		if e != nil {
			errors.Exit(e.Error())
		}

		o.Target = string(b)
	} else if o.URL != "" {

		o.Target = o.URL
	} else if o.List != "" {
		f, e := ioutil.ReadFile(o.List)
		if e != nil {
			errors.Exit(e.Error())
		}
		log.Println(string(f))
		o.Target = string(f)
	} else if o.payload || o.listApp || o.help || o.Version {
		// bypass vlidate target -flag payload
	} else {
		errors.Exit("No target input provided.")
	}

	if o.infosec_90 != "" {
		f, e := os.OpenFile(o.infosec_90,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if e != nil {
			errors.Exit(e.Error())
		}
		o.Output = f
	}
}

func isStdin() bool {
	f, e := os.Stdin.Stat()
	if e != nil {
		return false
	}
	if f.Mode()&os.ModeNamedPipe == 0 {
		return false
	}

	return true
}

func AppNamevalidate(flag string) bool {
	switch flag {
	case "salesforce":
		return true
	case "-payload", "--payload":
		return true
	}

	return false
}
