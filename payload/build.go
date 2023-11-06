package payload

import (
	"github.com/openebs/lib-csi/pkg/common/errors"

	"github.com/openebs/go-ogle-analytics/event"
)

type PayloadOption func(*Payload) error

type Payload struct {
	ClientId string     `json:"client_id"`
	Events   []ApiEvent `json:"events"`
}

type ApiEvent struct {
	Name   string             `json:"name"`
	Params event.OpenebsEvent `json:"params"`
}

func NewPayload(opts ...PayloadOption) (*Payload, error) {
	p := &Payload{}

	var err error
	for _, opt := range opts {
		err = opt(p)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build Payload")
		}
	}

	return p, nil
}

func WithClientId(clientId string) PayloadOption {
	return func(p *Payload) error {
		if len(clientId) == 0 {
			return errors.Errorf("failed to set Payload clientId: id is an empty string")
		}

		p.ClientId = clientId
		return nil
	}
}

func WithOpenebsEvent(event *event.OpenebsEvent) PayloadOption {
	return func(p *Payload) error {
		p.Events = append(p.Events, ApiEvent{
			Name:   event.CategoryStr() + "_" + event.ActionStr(),
			Params: *event,
		})
		return nil
	}
}
