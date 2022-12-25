# Eventt
Eventt is a small library to receive events/triggers from [Sonarr](https://github.com/Sonarr/Sonarr) using Webhook connections.

## What is the purpose of this library
Most of the tools/library communicate with Sonarr through the API, which is fine in a lot of cases, but sometimes we need a way to trigger action based on events happening in Sonarr rather than spamming the API every few secondes. Fortunately, Sonarr already have this mechanism implemented which called Webhook, and it has been used to send notification to other platforms like Discord or Slack.

If you're looking for a way to trigger action based on event on Sonarr this library is for you. otherwise if you want to interact with Sonarr like adding/deleting new shows or other functions, I recommend other libraries like [starr](https://github.com/golift/starr).

## Events
All the events/triggers is from the [Sonarr wiki](https://wiki.servarr.com/sonarr/settings#connection-triggers) and [webhook source code](https://github.com/Sonarr/Sonarr/blob/v3.0.9.1549/src/NzbDrone.Core/Notifications/Webhook/Webhook.cs):

| Event                         | Description |
| ----------------------------- | ----------------------------------------------------------------------------------------- |
| **OnGrab**                    | notified when episodes are available for download and has been sent to a download client. |
| **OnDownload**                | notified when episodes are successfully imported.                                         |
| **OnRename**                  | notified when episodes are renamed.                                                       |
| **OnSeriesDelete**            | notified when series are deleted.                                                         |
| **OnEpisodeFileDelete**       | notified when episodes files are deleted.                                                 |
| **OnHealth**                  | notified on health check failures.                                                        |
| **OnApplicationUpdate**       | notified when Sonarr gets updated to a new versions.                                      |
| **OnTest**                    | notified when test payload received.                                                      |
| **onUnknown**                 | notified when not implemented or unknown event received..                                 |

Note: until now there are no official documentation from Sonarr for webhook events JSON schema, therefore the current implementation for Go structure is based on running the service for a long time and collect payloads, then use it to restructure events body, if there is an issue with it or improvements please open an issue or send pull request and provide the payload from webhook event.

## Install

```shell
go get github.com/k-x7/eventt
```

## Usage
The following example demonstrate how to use **Eventt**:

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/k-x7/eventt"
)

func main() {

	events := eventt.SonarrTriggers{
		// Log on errors
		LogOnError: true,

		// on grab event print the show name
		OnGrab: func(event eventt.GrabEvent) {
			fmt.Printf("[Grab]: show name: %s\n", event.Series.Title)
		},

		// on download event print the show name
		OnDownload: func(event eventt.DownloadEvent) {
			fmt.Printf("[Download]: show name: %s\n", event.Series.Title)
		},

		// on test event print the show name
		OnTest: func(event eventt.TestEvent) {
			fmt.Printf("[Test]: sonarr send test event\n")
		},

		// if unknown event sent from sonarr print event type and payload
		OnUnknown: func(eventType string, event eventt.UnknownEvent) {
			fmt.Printf("[Unknown]: event type %s: %v\n", eventType, event)
		},
		// if on rename event sent from sonarr ignore it, this is the default action for all handlers if not set.
		OnRename: nil,

        // if any error happened while processing any request, print the error and payload, then send bad request http code for sonarr.
		OnError: func(payload []byte, err error) (httpStatus int) {
			fmt.Printf("[Error]: error: %v, payload: %v\n", err, payload)
			return http.StatusBadRequest
		},
	}

    // events will be received on http://localhost:8281/events
	http.HandleFunc("/events", events.Monitor)
	http.ListenAndServe("localhost:8281", nil)
}
```

then you can run it using `go run`:

```shell
$ go run main.go
```

Now we need to set Sonarr to send webhook events to this service, Go to your Sonarr webpage:

- Go to: **Settings** -> **Connect** -> **Click on Plus Sign** -> **Webhook**
- Add a **Name** for this connection.
- Select type of notification in **Notification Triggers** which you need to receive from Sonarr.
- Add **Tags** to limit webhook event for specific series if needed.
- Enter **URL**: `http://localhost:8281/events` or equivalent url based on your http service
- **Method** is not important for us you can leave it on `POST`
- Currently we don't implement **Username/Password** therefore leave it empty.
- Then click `Test` button, it should have a green check `✅` this mean Sonarr can send events to your service successfully.
- Press `Save` button and you're done.

Example: [Sonarr Webhook Settings Example](res/webhook-example.png)

Output from our service:
```shell
[Test]: sonarr send test event
[Test]: sonarr send test event
[Grab]: show name: Mob Psycho 100
[Download]: show name: Mob Psycho 100
[Download]: show name: Mob Psycho 100
[Test]: sonarr send test event
[Test]: sonarr send test event
```

# Usage Examples:

- [alertt](https://github.com:k-x7/alertt.git): alert user when grab or download events triggered using native system notification.