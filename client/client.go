package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/openebs/lib-csi/pkg/common/errors"

	"github.com/openebs/go-ogle-analytics/event"
	"github.com/openebs/go-ogle-analytics/payload"
)

func (c *MeasurementClient) Copy() *MeasurementClient {
	cpy := *c
	return &cpy
}

func (c *MeasurementClient) addFields(v url.Values) {
	v.Add("api_secret", c.apiSecret)
	v.Add("measurement_id", c.measurementId)
}

func (c *MeasurementClient) Send(event *event.OpenebsEvent) error {

	client := c.Copy()

	dataPayload, err := payload.NewPayload(
		payload.WithClientId(client.clientId),
		payload.WithOpenebsEvent(event),
	)

	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(dataPayload)
	if err != nil {
		return err
	}

	gaUrl := "https://www.google-analytics.com/mp/collect"

	req, err := http.NewRequest("POST", gaUrl, bytes.NewReader(jsonData))
	v := req.URL.Query()
	client.addFields(v)
	req.URL.RawQuery = v.Encode()

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode/100 != 2 {
		return errors.Errorf("Rejected by Google with code %d", resp.StatusCode)
	}

	return nil
}
