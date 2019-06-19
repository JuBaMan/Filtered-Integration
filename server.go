// Webhooks URL on Workspace: Development, Channel: slack-app:
// https://hooks.slack.com/services/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

// Generate new publicly accessible link to slash commands by running "ngrok http 192.168.99.100:5999" in a separate terminal window

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	webhookURL := "https://hooks.slack.com/services/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"

	// Add function to get the message from a server.
	message := `"Second message of the app."`

	// Add a function to format json.
	jsonBody := `{"text": ` + message + `}`

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		spew.Dump(jsonBody)
		spew.Dump(webhookURL)

		response, err := http.Post(webhookURL, "application/json", strings.NewReader(jsonBody))
		if err != nil {
			fmt.Println(err.Error())
		}
		defer response.Body.Close()
	})

	http.HandleFunc("/answer", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Post(webhookURL, "application/json", strings.NewReader(jsonBody))
		if err != nil {
			fmt.Println(err.Error())
		}
		defer response.Body.Close()
	})

	fmt.Println("Starting server...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
