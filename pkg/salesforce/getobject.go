package salesforce

import (
	"encoding/json"
	"log"
	"net/url"
	"regexp"

	"github.com/buger/jsonparser"
	"github.com/projectdiscovery/gologger"
)


func GetObjectList(ResponseGET []byte, url string, foundEndPoint []string, requestProxy string, requestHeaders []string) []string {
	var arrayObjList []string

	payloadMassage := string(PayloadGeneratorObjectList())
	fwuid, app, markup := GetAuraContext(ResponseGET)
	payloadAuraContext := string(PayloadGeneratorAuraContext(fwuid, app, markup))
	requestMethod := "POST"
	requestParameter := map[string]string{"message": payloadMassage, "aura.context": payloadAuraContext, "aura.token": "null"}
	responsebyte := RequestSalesforcePOST(url+foundEndPoint[0], requestMethod, requestProxy, requestHeaders, requestParameter)
	jsonparser.ArrayEach(responsebyte, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		jsonvalu, _, _, err := jsonparser.Get(value, "returnValue", "apiNamesToKeyPrefixes")
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
			gologger.Fatal().Msg("Failed to get ObjectList")
		}
		addValu := map[string]string{}
		json.Unmarshal([]byte(jsonvalu), &addValu)
		for k, _ := range addValu {
			arrayObjList = append(arrayObjList, k)
		}
	}, "actions")
	return arrayObjList
}
