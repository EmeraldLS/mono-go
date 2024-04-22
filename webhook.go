package monogo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/EmeraldLS/monogo/types"
	"github.com/EmeraldLS/monogo/utils"
)

func (w *MonoClient) SendWebhook(data *types.WebhookPayload) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return utils.MarshalError(err)
	}

	buf := bytes.NewBuffer(jsonBytes)
	w.withBody(buf)

	res, err := w.client.Do(w.req)
	if err != nil {
		return utils.RequestError(err)
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println("Error closing request body: ", err)
		}
	}(res.Body)

	status := "failed"

	if res.StatusCode == http.StatusOK {
		status = "delivered"
	}

	if status == "failed" {
		return errors.New(status)
	}

	log.Println("Status = ", status)

	return nil
}

func (m *MonoClient) ProcessWebhookRequest(ctx context.Context, secretKey string, bufferSize int, webhookQueue chan *types.WebhookPayload) {
	for payload := range webhookQueue {
		var whPayload = payload

		go func(*types.WebhookPayload) {
			backOffTime := time.Second
			maxBackOffTime := time.Hour
			retries := 0
			maxRetries := 5

			for {
				if err := m.SendWebhook(whPayload); err == nil {
					log.Println("Event sent successfully")
					break
				}

				retries++
				if retries >= maxRetries {
					log.Println("Event couldn't be completed")
					break
				}

				time.Sleep(backOffTime)
				backOffTime *= 2

				if backOffTime > maxBackOffTime {
					backOffTime = maxBackOffTime
				}
			}
		}(payload)
	}
}
