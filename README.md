## Go-ogle Analytics

Track and monitor your Go programs for free with Google Analytics

The `ga` package is essentially a Go wrapper around the [Google Analytics - Measurement Protocol (Google Analytics 4)](https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference?client_type=gtag)

### Install

```
go get -v github.com/openebs/go-ogle-analytics
```

### API

Create a new `client` and `Send()` an 'event'.

### Quick Usage

1. Log into GA and create a new property and note its Measurement ID

2. Create a `ga-test.go` file

	``` go
	package main

	import (
        gaClient "github.com/openebs/go-ogle-analytics/client"
        gaEvent "github.com/openebs/go-ogle-analytics/event"
	)

	func main() {
	    client, err := gaClient.NewMeasurementClient(
            gaClient.WithApiSecret("yourApiSecret"),
            gaClient.WithMeasurementId("G-yourMeasurementClient"),
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

	```

3. In GA, go to Report > Realtime

4. Run `ga-test.go`

	```
	$ go run ga-test.go
	Event fired!
	```

5. Watch as your event appears

	![foo-ga](https://cloud.githubusercontent.com/assets/633843/5979585/023fc580-a8fd-11e4-803a-956610bcc2e2.png)

#### MIT License

Copyright © 2015 &lt;dev@jpillora.com&gt;

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.