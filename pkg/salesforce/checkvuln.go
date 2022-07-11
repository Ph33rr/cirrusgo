package salesforce

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

func CheckVulnEndpoint(url string, requestProxy string, requestHeaders []string) []string {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	auraPathEndpoint := [7]string{
		"",
		"/", "/aura", "/s/aura", "/s/sfsites/aura", "/sfsites/aura", "/s/fact"}
	var foundEndPoint []string
	payloadPOST := string(PayloadGeneratorObjectList())
	requestMethod := "POST"
	requestParameter := map[string]string{"Massage": payloadPOST}

	for _, v := range auraPathEndpoint {

		responsebyte := RequestSalesforcePOST(url+v, requestMethod, requestProxy, requestHeaders, requestParameter)
		responsestr := string(responsebyte)
		if strings.Contains(responsestr, "aura:invalidSession") {
			foundEndPoint = append(foundEndPoint, v)

			fmt.Printf("[%s] %s %s\n", aurora.Red("VLN").String(),
				aurora.White(url+v).String(), aurora.Red("Vulnerable").String())

		}

	}
	if len(foundEndPoint) == 0 {
		fmt.Printf("[%s] %s %s\n", aurora.Blue("INFO").String(),
			aurora.White(url).String(), aurora.Blue("Not Vulnerable").String())
	} else {

		v := strings.Join(foundEndPoint, ",")
		fmt.Printf("[%s] %s\n", aurora.Red("vulnerable Endpoint").String(),
			aurora.White("["+v+"]").String())

	}

	return foundEndPoint
}
