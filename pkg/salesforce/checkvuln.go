package salesforce

import (
	"strings"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

func CheckVulnEndpoint(url string, requestProxy string, requestHeaders []string) []string {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	auraPathEndpoint := [10]string{
		"", "aura", "s/aura", "s/sfsites/aura", "sfsites/aura",
		"/", "/aura", "/s/aura", "/s/sfsites/aura", "/sfsites/aura"}
	var foundEndPoint []string
	payloadPOST := string(PayloadGeneratorObjectList())
	requestMethod := "POST"
	requestParameter := map[string]string{"Massage": payloadPOST}

	for _, v := range auraPathEndpoint {

		responsebyte := RequestSalesforcePOST(url+"/"+v, requestMethod, requestProxy, requestHeaders, requestParameter)
		responsestr := string(responsebyte)
		if strings.Contains(responsestr, "aura:invalidSession") {
			foundEndPoint = append(foundEndPoint, v)

			gologger.Warning().Msg("Website: " + url + " vulnerable Endpoint :" + v)
		}

	}
	if len(foundEndPoint) == 0 {
		gologger.Info().Msg("This Website: " + url + " Not vulnerable")
	} else {

		v := strings.Join(foundEndPoint, ",")
		gologger.Warning().Msg("Website: " + url + " vulnerable Endpoint: " + "[" + v + "]")
	}

	return foundEndPoint
}
