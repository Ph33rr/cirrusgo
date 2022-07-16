package salesforce

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/Ph33rr/cirrusgo/pkg/request"
	"github.com/projectdiscovery/gologger"
)

type dataSalesforce struct {
	requestURL       string
	requestMethod    string
	requestProxy     string
	requestHeaders   []string
	requestParameter map[string]string
}

func RequestSalesforcePOST(requestURL string, requestMethod string, requestProxy string, requestHeaders []string, requestParameter map[string]string) []byte {

	client := request.Client(requestProxy)
	requestParam := url.Values{}

	//add all param
	if requestParameter != nil {

		for keyParameter, valuParameter := range requestParameter {
			requestParam.Set(keyParameter, valuParameter)
		}
	}
	//encode all param
	finalRequestBody := bytes.NewBuffer([]byte(requestParam.Encode()))
	request, err := http.NewRequest(requestMethod, requestURL, finalRequestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {

		log.Fatalf("An Error Occured %v", err)
		gologger.Fatal().Msg("can't handle Add Header(POST)")

	}

	//add all header
	if requestHeaders != nil {
		for _, header := range requestHeaders {
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
		gologger.Fatal().Msg("Can't send POST request")

	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
		gologger.Fatal().Msg("Can't handle response(POST)")
	}

	defer request.Body.Close()
	//return respnse
	return responseBody
}

//this func return response (fwuid and markup for salseforce)
func RequestSalesforceGET(requestURL string, requestMethod string, requestProxy string, requestHeaders []string, requestParameter map[string]string) []byte {

	client := request.Client(requestProxy)
	requestParam := url.Values{}
	//add param
	if requestParameter != nil {

		for keyParameter, valuParameter := range requestParameter {

			requestParam.Set(keyParameter, valuParameter)
		}
	}
	//encode param
	finalRequestBody := bytes.NewBuffer([]byte(requestParam.Encode()))
	request, err := http.NewRequest(requestMethod, requestURL, finalRequestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		gologger.Fatal().Msg("can't handle Add Header(GET)")
	}
	//add header

	if requestHeaders != nil {
		for _, header := range requestHeaders {
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

		gologger.Fatal().Msg("Can't send GET request")
	}

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
		gologger.Fatal().Msg("Can't handle response (GET)")
	}

	//check header if respnse.body == null follow redircet

	checkResponseBody := string(responseBody)
	//bytes.Equal(responseBody, emptyByteVar)
	// not working for all website
	//emptyByteVar := make([]byte, 128)

	if strings.Contains(checkResponseBody, "fwuid") {
		return responseBody

	} else if len(checkResponseBody) == 0 {
		// get url from Location header

		responseHeader := response.Header.Get("Location")
		requestURL := responseHeader
		responseBody := SalesforceGetURLFromBody(requestURL, requestMethod, requestProxy, requestHeaders, requestParameter)
		checkResponseBody := string(responseBody)

		if strings.Contains(checkResponseBody, "fwuid") {
			return responseBody
		} else {
			checkResponseBody := string(responseBody)
			if strings.Contains(checkResponseBody, "window.location.href ='") {

				re := regexp.MustCompile("window.location.href ='([^']+)")
				getURLFromBody := re.FindString(checkResponseBody)
				responseHeader := string(getURLFromBody[23:])
				requestURL := responseHeader
				responseBody := SalesforceGetURLFromBody(requestURL, requestMethod, requestProxy, requestHeaders, requestParameter)

				if strings.Contains(checkResponseBody, "fwuid") {

					return responseBody

				} else {

					responseHeader := SalesforceGetURLFromHeader(requestURL, requestMethod, requestProxy, requestHeaders, requestParameter)
					requestURL := responseHeader
					responseBody := SalesforceGetURLFromBody(requestURL, requestMethod, requestProxy, requestHeaders, requestParameter)
					checkResponseBody := string(responseBody)
					if strings.Contains(checkResponseBody, "fwuid") {

						return responseBody
					}
				}
			}
		}
	} else if strings.Contains(checkResponseBody, "window.location.href ='") {
		re := regexp.MustCompile("window.location.href ='([^']+)")
		getURLFromBody := re.FindString(checkResponseBody)
		if strings.Contains(getURLFromBody, "http") {
			responseHeader := string(getURLFromBody[23:])
			requestURL := responseHeader
			responseBody := SalesforceGetURLFromBody(requestURL, requestMethod, requestProxy, requestHeaders, requestParameter)
			checkResponseBody := string(responseBody)
			if strings.Contains(checkResponseBody, "fwuid") {
				return responseBody
			}
		}
		if !strings.Contains(getURLFromBody, "http") {
			responseHeader := string(getURLFromBody[23:])
			requestURL := requestURL + responseHeader
			responseBody := SalesforceGetURLFromBody(requestURL, requestMethod, requestProxy, requestHeaders, requestParameter)
			checkResponseBody := string(responseBody)
			if strings.Contains(checkResponseBody, "fwuid") {
				return responseBody
			} else {
				responseHeader := SalesforceGetURLFromHeader(requestURL, requestMethod, requestProxy, requestHeaders, requestParameter)
				requestURL := responseHeader
				responseBody := SalesforceGetURLFromBody(requestURL, requestMethod, requestProxy, requestHeaders, requestParameter)
				checkResponseBody := string(responseBody)
				if strings.Contains(checkResponseBody, "fwuid") {

					return responseBody
				}
			}
		} else {
			responseHeader := SalesforceGetURLFromHeader(requestURL, requestMethod, requestProxy, requestHeaders, requestParameter)
			requestURL := responseHeader
			responseBody := SalesforceGetURLFromBody(requestURL, requestMethod, requestProxy, requestHeaders, requestParameter)
			checkResponseBody := string(responseBody)
			if strings.Contains(checkResponseBody, "fwuid") {

				return responseBody
			}
		}
	}

	defer request.Body.Close()
	return responseBody
}

//return request
func SalesforceGetURLFromBody(requestURL string, requestMethod string, requestProxy string, requestHeaders []string, requestParameter map[string]string) []byte {

	client := request.Client(requestProxy)
	requestParam := url.Values{}

	if requestParameter != nil {
		for keyParameter, valuParameter := range requestParameter {

			requestParam.Set(keyParameter, valuParameter)
		}
	}
	//encode all param
	finalRequestBody := bytes.NewBuffer([]byte(requestParam.Encode()))
	request, err := http.NewRequest(requestMethod, requestURL, finalRequestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	if requestHeaders != nil {

		for _, header := range requestHeaders {
			parts := strings.SplitN(header, ":", 2)
			if len(parts) != 2 {
				continue
			}

			request.Header.Set(parts[0], parts[1])
		}
	}

	response, err := client.Do(request)

	if err != nil {
		gologger.Fatal().Msg("Can't GET AppName,UID,Markup try manual with flag `cirrusgo salesforce -payload -help`")

	}

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
		gologger.Fatal().Msg("Can't handle response (fromBody)")
	}

	defer request.Body.Close()
	return responseBody
}

func SalesforceGetURLFromHeader(requestURL string, requestMethod string, requestProxy string, requestHeaders []string, requestParameter map[string]string) string {

	client := request.Client(requestProxy)
	requestParam := url.Values{}

	if requestParameter != nil {
		for keyParameter, valuParameter := range requestParameter {

			requestParam.Set(keyParameter, valuParameter)
		}
	}
	//encode all param
	finalRequestBody := bytes.NewBuffer([]byte(requestParam.Encode()))
	request, err := http.NewRequest(requestMethod, requestURL, finalRequestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		gologger.Fatal().Msg("Can't send GET request (fromheader)")
	}

	if requestHeaders != nil {
		for _, header := range requestHeaders {
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

	responseHeader := response.Header.Get("Location")
	if err != nil {
		log.Fatalln(err)
		gologger.Fatal().Msg("Can't handle response (fromHeader)")

	}

	defer request.Body.Close()
	return responseHeader
}
