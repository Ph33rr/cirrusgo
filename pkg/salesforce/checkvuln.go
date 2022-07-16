package salesforce

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

func CheckVulnEndpoint(urls string, requestProxy string, requestHeaders []string) ([]string, string) {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	domainfromURL := urls
	auraPathEndpoint := [7]string{
		"",
		"/", "/aura", "/s/aura", "/s/sfsites/aura", "/sfsites/aura", "/s/fact"}
	var foundEndPoint []string
	payloadPOST := string(PayloadGeneratorObjectList())
	requestMethod := "POST"
	requestParameter := map[string]string{"Massage": payloadPOST}

	for _, v := range auraPathEndpoint {

		responsebyte := RequestSalesforcePOST(urls+v, requestMethod, requestProxy, requestHeaders, requestParameter)
		responsestr := string(responsebyte)
		if strings.Contains(responsestr, "aura:invalidSession") {
			foundEndPoint = append(foundEndPoint, v)

			fmt.Printf("[%s] %s %s\n", aurora.Red("VLN").String(),
				aurora.White(urls+v).String(), aurora.Red("Vulnerable").String())

		}

	}
	if len(foundEndPoint) == 0 {
		urlso, err := url.Parse(domainfromURL)
		if err != nil {
			log.Fatal(err)
		}
		hostname := strings.TrimPrefix(urlso.Hostname(), "")
		hostname2 := strings.TrimPrefix(urlso.Scheme, "")
		urls := hostname2 + "://" + hostname
		for _, v := range auraPathEndpoint {

			responsebyte := RequestSalesforcePOST(urls+v, requestMethod, requestProxy, requestHeaders, requestParameter)
			responsestr := string(responsebyte)
			if strings.Contains(responsestr, "aura:invalidSession") {
				foundEndPoint = append(foundEndPoint, v)

				fmt.Printf("[%s] %s %s\n", aurora.Red("VLN").String(),
					aurora.White(urls+v).String(), aurora.Red("Vulnerable").String())

			}
		}
		if len(foundEndPoint) == 0 {
			fmt.Printf("[%s] %s %s\n", aurora.Blue("INFO").String(),
				aurora.White(urls).String(), aurora.Blue("Not Vulnerable").String())
		} else {
			v := strings.Join(foundEndPoint, ",")
			fmt.Printf("[%s] %s\n", aurora.Red("vulnerable Endpoint").String(),
				aurora.White("["+v+"]").String())
			return foundEndPoint, urls
		}

	} else {

		v := strings.Join(foundEndPoint, ",")
		fmt.Printf("[%s] %s\n", aurora.Red("vulnerable Endpoint").String(),
			aurora.White("["+v+"]").String())

	}

	return foundEndPoint, urls
}
