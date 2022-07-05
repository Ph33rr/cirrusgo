package cirrusgo

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Ph33rr/CirrusGo/pkg/request"
)


func (dataRequest dataSalesforce) requestSalesforcePOST() []byte {

	client := request.Client(dataRequest.requestProxy)
	requestParam := url.Values{}

	if dataRequest.requestParameter != nil {
		for keyParameter, valuParameter := range dataRequest.requestParameter {

			requestParam.Set(keyParameter, valuParameter)
		}
	}

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
	//return request
	return responseBody
}
