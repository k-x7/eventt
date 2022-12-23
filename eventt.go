package eventt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/exp/slog"
)

// SonarrTriggers sonarr events or triggers using webhook connection
// see: https://wiki.servarr.com/sonarr/settings#connection-triggers
type SonarrTriggers struct {
	// OnGrab notified when episodes are available for download and has been sent to a download client
	OnGrab func(event GrabEvent)
	// OnDownload or OnImport be notified when episodes are successfully imported
	OnDownload func(event DownloadEvent)
	// OnRename be notified when episodes are renamed
	OnRename func(event RenameEvent)
	// OnEpisodeFileDelete be notified when episodes files are deleted
	OnEpisodeFileDelete func(event EpisodeFileDeleteEvent)
	// OnSeriesDelete be notified when series are deleted
	OnSeriesDelete func(event SeriesDeleteEvent)
	// OnHealth be notified on health check failures
	OnHealth func(event HealthEvent)
	// OnApplicationUpdate be notified when Sonarr gets updated to a new version
	OnApplicationUpdate func(event ApplicationUpdateEvent)
	// OnTest be notified when test payload received
	OnTest func(event TestEvent)
	// OnUnknown be notified when not implemented or unknown event received.
	OnUnknown func(eventType string, event UnknownEvent)
	// OnError callback for any error process any event
	// the payload represents the received request it can be nil if the error
	// reading the payload, other error will include the payload
	// this function should return http status code to return to Sonarr.
	// if this function not implemented, it will log any error and return 400 to sonarr
	OnError func(payload []byte, err error) (httpStatus int)
	// LogOnError should we log errors, if true it will use slog.Error to log errors and
	// it will include the payload. see SonarrTriggers.handleErrors for more details.
	LogOnError bool
	b          []byte
}

// Monitor http handler to invoke the correct trigger from SonarrTriggers based on
// the received event from Sonarr, see SonarrTriggers all the events and error handling.
func (s *SonarrTriggers) Monitor(w http.ResponseWriter, r *http.Request) {
	var err error

	s.b, err = io.ReadAll(r.Body)
	if err != nil {
		status := s.handleErrors(
			fmt.Errorf("error while reading request body %w", err),
		)
		w.WriteHeader(status)
		return
	}
	r.Body.Close()

	eventType := &WebhookEvent{}

	if err := json.Unmarshal(s.b, eventType); err != nil {
		status := s.handleErrors(
			fmt.Errorf("error parsing event type: %w", err),
		)
		w.WriteHeader(status)
		return
	}

	if err := s.handleEvent(eventType.EventType); err != nil {
		status := s.handleErrors(
			fmt.Errorf("error handle '%s' event: %w", eventType.EventType, err),
		)
		w.WriteHeader(status)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *SonarrTriggers) handleEvent(eventType string) error {
	switch eventType {
	case grab:
		return handleGenericEvent(s.b, s.OnGrab)
	case download:
		return handleGenericEvent(s.b, s.OnDownload)
	case rename:
		return handleGenericEvent(s.b, s.OnRename)
	case episodeFileDelete:
		return handleGenericEvent(s.b, s.OnEpisodeFileDelete)
	case seriesDelete:
		return handleGenericEvent(s.b, s.OnSeriesDelete)
	case health:
		return handleGenericEvent(s.b, s.OnHealth)
	case applicationUpdate:
		return handleGenericEvent(s.b, s.OnApplicationUpdate)
	case test:
		return handleGenericEvent(s.b, s.OnTest)
	default:
		return s.handleUnknown(eventType)
	}
}

func (s *SonarrTriggers) handleErrors(err error) int {
	status := http.StatusBadRequest
	if s.OnError != nil {
		status = s.OnError(s.b, err)
	}
	if s.LogOnError {
		slog.Error("error processing new event", err, "payload", string(s.b))
	}
	return status
}

func handleGenericEvent[T eventType](b []byte, f func(e T)) error {
	if f == nil {
		return nil
	}
	var e T
	if err := json.Unmarshal(b, &e); err != nil {
		return fmt.Errorf("error parsing '%s' event: %w", e.eventName(), err)
	}
	f(e)
	return nil
}

func (s *SonarrTriggers) handleUnknown(eventType string) error {
	if s.OnUnknown == nil {
		return nil
	}
	m := make(UnknownEvent)
	if err := json.Unmarshal(s.b, &m); err != nil {
		return fmt.Errorf("error parsing 'Unknown' event: %w", err)
	}
	s.OnUnknown(eventType, m)
	return nil
}
