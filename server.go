// Webhooks URL on Workspace: Development, Channel: slack-app:
// https://hooks.slack.com/services/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

package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	webhookURL := "https://hooks.slack.com/services/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"
	jsonBody := new(bytes.Buffer)

	// Add function to get the message from a server.
	message := `"Second message of the app."`

	// Add a function to format json.
	jsonBody.WriteString(`{"text":`)
	jsonBody.WriteString(message)
	jsonBody.WriteString(`}`)

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Post(webhookURL, "application/json", jsonBody)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer response.Body.Close()
		log.Println(response)
	})

	fmt.Println("Starting server...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
