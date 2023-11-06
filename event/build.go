package event

import (
	"github.com/openebs/lib-csi/pkg/common/errors"
)

type OpenebsEventOption func(*OpenebsEvent) error

// OpenebsEvent Hit Type
type OpenebsEvent struct {
	Category string `json:"category"`
	Action   string `json:"action"`
	Label    string `json:"label"`
	Value    int64  `json:"value"`
}

func NewOpenebsEvent(opts ...OpenebsEventOption) (*OpenebsEvent, error) {
	e := &OpenebsEvent{}

	var err error
	for _, opt := range opts {
		err = opt(e)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build OpenebsEvent")
		}
	}

	return e, nil
}

func WithCategory(category string) OpenebsEventOption {
	return func(e *OpenebsEvent) error {
		if len(category) == 0 {
			return errors.Errorf("failed to set OpenebsEvent category: category is an empty string")
		}

		e.Category = category
		return nil
	}
}

func WithAction(action string) OpenebsEventOption {
	return func(e *OpenebsEvent) error {
		if len(action) == 0 {
			return errors.Errorf("failed to set OpenebsEvent action: action is an empty string")
		}

		e.Action = action
		return nil
	}
}

func WithLabel(label string) OpenebsEventOption {
	return func(e *OpenebsEvent) error {
		if len(label) == 0 {
			return errors.Errorf("failed to set OpenebsEvent label: label is an empty string")
		}

		e.Label = label
		return nil
	}
}

func WithValue(value int64) OpenebsEventOption {
	return func(e *OpenebsEvent) error {
		e.Value = value
		return nil
	}
}
