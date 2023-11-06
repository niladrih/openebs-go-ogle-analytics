package main

import (
	gaClient "github.com/openebs/go-ogle-analytics/client"
	gaEvent "github.com/openebs/go-ogle-analytics/event"
)

func main() {
	client, err := gaClient.NewMeasurementClient(
		gaClient.WithApiSecret("NguBiGh6QeOdeG3zJswggQ"),
		gaClient.WithMeasurementId("G-TZGP46618W"),
		gaClient.WithClientId("uniqueUserId-000000001"),
	)
	if err != nil {
		panic(err)
	}

	event, err := gaEvent.NewOpenebsEvent(
		gaEvent.WithCategory("Foo"),
		gaEvent.WithAction("Bar"),
		gaEvent.WithLabel("Baz"),
		gaEvent.WithValue(19072023),
	)
	if err != nil {
		panic(err)
	}

	err = client.Send(event)
	if err != nil {
		panic(err)
	}

	println("Event fired!")
}
