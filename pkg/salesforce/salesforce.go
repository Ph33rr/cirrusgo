package salesforce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/ph33rr/CirrusGo/pkg/request"
)

type payload struct {
	Massage json.RawMessage `json:"actions"`
}
