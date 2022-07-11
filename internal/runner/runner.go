package runner

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
        "github.com/ph33rr/CirrusGo/pkg/salesforce"
)

func New(options *Options) {

	showBanner()

	if len(os.Args) != 1 && AppNamevalidate(os.Args[1]) {
		if o.payload {
			if o.ObjectList {
				fmt.Printf("[%s] %s %s\n", aurora.Green("Generator").String(),
					aurora.White("Payload:").String(), aurora.Blue(string(salesforce.PayloadGeneratorObjectList())))
			}
			if o.Dump {
				fmt.Printf("[%s] %s %s\n", aurora.Green("Generator").String(),
					aurora.White("Payload:").String(), aurora.Blue(string(salesforce.PayloadGeneratorDump())))
			}
			if o.GetItems {
				if o.objName != "" && o.pageNumber != 0 && o.pageSize != 0 {
					fmt.Printf("[%s] %s %s\n", aurora.Green("Generator").String(),
						aurora.White("Payload:").String(), aurora.Blue(string(salesforce.PayloadGeneratorGetItems(o.objName, o.pageSize, o.pageNumber))))
				}
			}
			if o.GetRecord {
				if o.recordId != "" {
					fmt.Printf("[%s] %s %s\n", aurora.Green("Generator").String(),
						aurora.White("Payload:").String(), aurora.Blue(string(salesforce.PayloadGeneratorGetRecord(o.recordId))))
				}
			}
			if o.WritableOBJ {
				if o.objName != "" {
					fmt.Printf("[%s] %s %s\n", aurora.Green("Generator").String(),
						aurora.White("Payload:").String(), aurora.Blue(string(salesforce.PayloadGeneratorWritableOBJ(o.objName))))
				}
			}
			if o.SearchObj {
				if o.objName != "" && o.pageNumber != 0 && o.pageSize != 0 {
					fmt.Printf("[%s] %s %s\n", aurora.Green("Generator").String(),
						aurora.White("Payload:").String(), aurora.Blue(string(salesforce.PayloadGeneratorSearchObj(o.objName, o.pageSize, o.pageNumber))))
				}
			}
			if o.AuraContext {
				if o.fwuid != "" && o.AppName != "" && o.markup != "" {
					fmt.Printf("[%s] %s %s\n", aurora.Green("Generator").String(),
						aurora.White("Payload:").String(), aurora.Blue(string(salesforce.PayloadGeneratorAuraContext(o.fwuid, o.AppName, o.markup))))
				}
			} else {
				fmt.Println(salesforceOptionsPayload)
			}
		} else {
			if o.check {
				salesforce.CheckVulnEndpoint(o.Target, o.Proxy, o.Headers)
			}
		}
		if o.listobj {
			foundEndPoint := salesforce.CheckVulnEndpoint(o.Target, o.Proxy, o.Headers)
			ResponseGET := salesforce.RequestSalesforceGET(o.Target, "GET", o.Proxy, o.Headers, nil)
			fmt.Printf("[%s] %s %s\n", aurora.Green("INFO").String(),
				aurora.White("Object List:").String(), aurora.White(salesforce.GetObjectList(ResponseGET, o.Target, foundEndPoint, o.Proxy, o.Headers)))
		}
		if o.getobj {
			foundEndPoint := salesforce.CheckVulnEndpoint(o.Target, o.Proxy, o.Headers)
			ResponseGET := salesforce.RequestSalesforceGET(o.Target, "GET", o.Proxy, o.Headers, nil)
			fmt.Printf("[%s] %s %s %s\n", aurora.Red("Exploit").String(),
				aurora.White("Object").String(), aurora.White(o.objName+":").String(),
				aurora.White(salesforce.GetObjectItems(ResponseGET, o.Target, foundEndPoint, o.objName, o.pageSize, o.pageNumber, o.Proxy, o.Headers)))

		}
		if o.getrecord {
			foundEndPoint := salesforce.CheckVulnEndpoint(o.Target, o.Proxy, o.Headers)
			ResponseGET := salesforce.RequestSalesforceGET(o.Target, "GET", o.Proxy, o.Headers, nil)
			fmt.Printf("[%s] %s %s %s\n", aurora.Red("Exploit").String(),
				aurora.White("Object").String(), aurora.White(o.objName+":").String(),
				aurora.White(salesforce.GetObjectRecord(ResponseGET, o.Target, foundEndPoint, o.recordId, o.Proxy, o.Headers)))

		}
		if o.full {
			foundEndPoint := salesforce.CheckVulnEndpoint(o.Target, o.Proxy, o.Headers)
			ResponseGET := salesforce.RequestSalesforceGET(o.Target, "GET", o.Proxy, o.Headers, nil)
			objectList := salesforce.GetObjectList(ResponseGET, o.Target, foundEndPoint, o.Proxy, o.Headers)
			for _, v := range objectList {
				response := salesforce.GetObjectItems(ResponseGET, o.Target, foundEndPoint, v, o.pageSize, o.pageNumber, o.Proxy, o.Headers)
				if len(response) <= 2 {
					fmt.Printf("[%s] %s %s \n", aurora.Green("INFO").String(),
						aurora.White("Object empty").String(), aurora.White(v+":"))
				} else {
					fmt.Printf("[%s] %s %s %s\n", aurora.Red("Exploit").String(),
						aurora.White("Object").String(), aurora.White(v+":").String(),
						aurora.White(response))
				}
			}
		}
		if o.checkWritableOBJ {
			foundEndPoint := salesforce.CheckVulnEndpoint(o.Target, o.Proxy, o.Headers)
			ResponseGET := salesforce.RequestSalesforceGET(o.Target, "GET", o.Proxy, o.Headers, nil)
			objectList := salesforce.GetObjectList(ResponseGET, o.Target, foundEndPoint, o.Proxy, o.Headers)
			for _, v := range objectList {
				response := salesforce.GetWritableObject(ResponseGET, o.Target, foundEndPoint, v, o.Proxy, o.Headers)
				if strings.Contains(response, "fields") {
					fmt.Printf("[%s] %s %s \n", aurora.Red("Exploit").String(),
						aurora.White("Object Writable").String(), aurora.White(v+":"))
				} else {
					fmt.Printf("[%s] %s %s\n", aurora.Green("INFO").String(),
						aurora.White("Object Not Writable").String(), aurora.White(v+":").String())
				}
			}

		}
		if o.dump {
			foundEndPoint := salesforce.CheckVulnEndpoint(o.Target, o.Proxy, o.Headers)
			ResponseGET := salesforce.RequestSalesforceGET(o.Target, "GET", o.Proxy, o.Headers, nil)
			fmt.Printf("[%s] %s %s\n", aurora.Green("INFO").String(),
				aurora.White("Dump all INFO:").String(), aurora.White(salesforce.GetDump(ResponseGET, o.Target, foundEndPoint, o.Proxy, o.Headers)))

		}
		if o.help {
			fmt.Println(salesforceOptions)
		}

	} else if isURL(FixURL(options.Target)) {
		salesforce.CheckVulnEndpoint(FixURL(options.Target), options.Proxy, options.Headers)
	}
	if o.listApp {
		fmt.Printf("\n[%s] %s\n", aurora.Green("App Support").String(),
			aurora.White(applist).String())

	}
	if o.help {
		h :=
			"\n" +
				"Usage:" + usage +
				"\n" +
				"Options:" + optionss +
				""

		fmt.Println(h)
	}
	if o.Version {
		showVersion()

	}
}
func isURL(s string) bool {

	_, e := url.ParseRequestURI(s)
	if e != nil {
		return false
	}

	u, e := url.Parse(s)
	log.Println(s)
	if e != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func FixURL(s string) string {

	oldurl := strings.ReplaceAll(s, "\r\n", "")
	newurl := strings.ReplaceAll(oldurl, "\ufeff", "")

	return newurl
}
