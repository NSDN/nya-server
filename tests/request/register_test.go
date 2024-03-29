package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/NSDN/nya-server/configs"
)

func TestRegisterRequest(t *testing.T) {
	data := map[string]string{"endpoint": "endpoint"}

	jsonData, err := json.Marshal(data)

	if err != nil {
		t.Error("Error encoding JSON: ", err)
		return
	}

	response, err := http.Post(configs.API_BASE_URL+configs.API_REGISTER, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		t.Error("Error making HTTP request: ", err)
		return
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Error("Error reading response body: ", err)
		return
	}

	t.Log(string(body))
}
