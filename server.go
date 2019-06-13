// Webhooks URL on Workspace: Development, Channel: slack-app:
// https://hooks.slack.com/services/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

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	webhookURL := "https://hooks.slack.com/services/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"
	jsonBody := new(bytes.Buffer)

	// The client should be integrated with a microservice running on a docker machine.
	client := &http.Client{}

	// Add function to get the message from a server.
	message := `"First message of the app."`

	// Add a function to format json.
	jsonBody.WriteString(`{"text":`)
	jsonBody.WriteString(message)
	jsonBody.WriteString(`}`)

	request, err := http.NewRequest("POST", webhookURL, jsonBody)
	if err != nil {
		panic(err)
	}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Diagnosis.
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
}
