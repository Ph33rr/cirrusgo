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


func GetObjectItems(ResponseGET []byte, url string, foundEndPoint []string, objectName string, pageSize int, page int, requestProxy string, requestHeaders []string) string {
	var data string
	payloadMassage := string(PayloadGeneratorGetItems(objectName, pageSize, page))
	fwuid, app, markup := GetAuraContext(ResponseGET)
	payloadAuraContext := string(PayloadGeneratorAuraContext(fwuid, app, markup))
	requestMethod := "POST"
	requestParameter := map[string]string{"message": payloadMassage, "aura.context": payloadAuraContext, "aura.token": "null"}

	responsebyte := RequestSalesforcePOST(url+foundEndPoint[0], requestMethod, requestProxy, requestHeaders, requestParameter)
	jsonparser.ArrayEach(responsebyte, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		jsonvalu, _, _, err := jsonparser.Get(value, "returnValue", "result")
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
			gologger.Fatal().Msg("Failed to get ObjectItems")
		}
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
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
			gologger.Fatal().Msg("Failed to get ObjectWritable")
		}
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

