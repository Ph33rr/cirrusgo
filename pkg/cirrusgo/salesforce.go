package cirrusgo

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/ph33rr/CirrusGo/pkg/request"
)


func (dataRequest dataSalesforce) requestSalesforcePOST() []byte {

	client := request.Client(dataRequest.requestProxy)
	requestParam := url.Values{}

	//add all param
	if dataRequest.requestParameter != nil {
		for keyParameter, valuParameter := range dataRequest.requestParameter {

			requestParam.Set(keyParameter, valuParameter)
		}
	}
	//encode all param
	finalRequestBody := bytes.NewBuffer([]byte(requestParam.Encode()))

	request, err := http.NewRequest(dataRequest.requestMethod, dataRequest.requestURL, finalRequestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	//add all header
	if dataRequest.requestHeaders != nil {
		for _, header := range dataRequest.requestHeaders {
			parts := strings.SplitN(header, ":", 2)

			if len(parts) != 2 {
				continue
			}

			request.Header.Set(parts[0], parts[1])
		}
	}
	//send request
	response, err := client.Do(request)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	defer request.Body.Close()
	//return respnse
	return responseBody
}

//this func return response (fwuid and markup for salseforce)
func (dataRequest dataSalesforce) requestSalesforceGET() []byte {

	client := request.Client(dataRequest.requestProxy)
	requestParam := url.Values{}
	//add param
	if dataRequest.requestParameter != nil {
		for keyParameter, valuParameter := range dataRequest.requestParameter {

			requestParam.Set(keyParameter, valuParameter)
		}
	}
	//encode param
	finalRequestBody := bytes.NewBuffer([]byte(requestParam.Encode()))

	request, err := http.NewRequest(dataRequest.requestMethod, dataRequest.requestURL, finalRequestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	//add header
	if dataRequest.requestHeaders != nil {
		for _, header := range dataRequest.requestHeaders {
			parts := strings.SplitN(header, ":", 2)

			if len(parts) != 2 {
				continue
			}

			request.Header.Set(parts[0], parts[1])
		}
	}
	//send request
	response, err := client.Do(request)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)

	}

	//check header if respnse.body == null follow redircet

	checkResponseBody := string(responseBody)
	//bytes.Equal(responseBody, emptyByteVar)
	// not working for all website
	//emptyByteVar := make([]byte, 128)

	if strings.Contains(checkResponseBody, "fwuid") {

		return responseBody

	} else if strings.Contains(checkResponseBody, "") {
		// get url from Location header
		responseHeader := response.Header.Get("Location")
		dataRequest.requestURL = responseHeader
		responseBody = dataRequest.SalesforceGetURLFromBody()

		if strings.Contains(checkResponseBody, "fwuid") {
			return responseBody

		} else {
			checkResponseBody := string(responseBody)
			if strings.Contains(checkResponseBody, "window.location.href ='") {

				re := regexp.MustCompile("window.location.href ='([^']+)")
				getURLFromBody := re.FindString(checkResponseBody)
				responseHeader := string(getURLFromBody[23:])
				dataRequest.requestURL = responseHeader
				responseBody = dataRequest.SalesforceGetURLFromBody()

				if strings.Contains(checkResponseBody, "fwuid") {

					return responseBody

				} else {
					dataRequest.requestURL = responseHeader
					responseHeader := dataRequest.SalesforceGetURLFromHeader()
					dataRequest.requestURL = responseHeader
					responseBody = dataRequest.SalesforceGetURLFromBody()

					if strings.Contains(checkResponseBody, "fwuid") {

						return responseBody
					}
				}
			}
		}
	} else if strings.Contains(checkResponseBody, "window.location.href ='") {

		re := regexp.MustCompile("window.location.href ='([^']+)")
		getURLFromBody := re.FindString(checkResponseBody)
		responseHeader := string(getURLFromBody[23:])
		dataRequest.requestURL = responseHeader
		responseBody = dataRequest.SalesforceGetURLFromBody()

		if strings.Contains(checkResponseBody, "fwuid") {

			return responseBody
		}
	}

	defer request.Body.Close()
	return responseBody
}

//return request
func (dataRequest dataSalesforce) SalesforceGetURLFromBody() []byte {

	client := request.Client(dataRequest.requestProxy)
	requestParam := url.Values{}

	if dataRequest.requestParameter != nil {
		for keyParameter, valuParameter := range dataRequest.requestParameter {

			requestParam.Set(keyParameter, valuParameter)
		}
	}
	//encode all param
	finalRequestBody := bytes.NewBuffer([]byte(requestParam.Encode()))

	request, err := http.NewRequest(dataRequest.requestMethod, dataRequest.requestURL, finalRequestBody)

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	if dataRequest.requestHeaders != nil {

		for _, header := range dataRequest.requestHeaders {
			parts := strings.SplitN(header, ":", 2)

			if len(parts) != 2 {
				continue
			}

			request.Header.Set(parts[0], parts[1])
		}
	}

	response, err := client.Do(request)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)

	}

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	defer request.Body.Close()
	return responseBody
}

func (dataRequest dataSalesforce) SalesforceGetURLFromHeader() string {

	client := request.Client(dataRequest.requestProxy)
	requestParam := url.Values{}

	if dataRequest.requestParameter != nil {
		for keyParameter, valuParameter := range dataRequest.requestParameter {

			requestParam.Set(keyParameter, valuParameter)
		}
	}
	//encode all param
	finalRequestBody := bytes.NewBuffer([]byte(requestParam.Encode()))

	request, err := http.NewRequest(dataRequest.requestMethod, dataRequest.requestURL, finalRequestBody)

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	if dataRequest.requestHeaders != nil {
		
		for _, header := range dataRequest.requestHeaders {
			parts := strings.SplitN(header, ":", 2)

			if len(parts) != 2 {
				continue
			}

			request.Header.Set(parts[0], parts[1])
		}
	}

	response, err := client.Do(request)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)

	}

	responseBody := response.Header.Get("Location")

	if err != nil {
		log.Fatalln(err)
	}

	defer request.Body.Close()
	return responseBody
}
