package request

import (
	"io"
	"net/http"
	"testing"

	"github.com/NSDN/nya-server/configs"
	"github.com/NSDN/nya-server/utils"
)

func TestPingRequest(t *testing.T) {
	response, err := http.Get(configs.API_BASE_URL + "/ping")

	if err != nil {
		utils.HandlePrintExample("Error making HTTP request: ", err)
		return
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		utils.HandlePrintExample("Error reading response body: ", err)
		return
	}

	t.Logf("Body: %s", string(body))
}
