package publisher

import (
	"falabella.com/boss-message-library/internal/formatter"
	"github.com/google/uuid"
)

const (
	COMMERCE   = "FALABELLA"
	VERSION    = "1.0"
	CHANNEL    = "WEB"
	CAPABILITY = "capa"
	DOMAIN     = "FAL"
	MIMETYPE   = "application/json"
)

func CompleteHeader(attributes map[string]string, eventType, entityType string) {
	if attributes == nil {
		attributes = make(map[string]string)
	}
	timeStamp := formatter.Timestamp()
	dateTime := formatter.LocalDateTimeUTC()
	if attributes["eventId"] == "" {
		attributes["eventId"] = uuid.NewString()
	}
	if attributes["entityId"] == "" {
		attributes["entityId"] = uuid.NewString()
	}
	attributes["entityType"] = entityType
	attributes["eventType"] = eventType
	attributes["datetime"] = dateTime
	attributes["timestamp"] = timeStamp
	if attributes["version"] == "" {
		attributes["version"] = VERSION
	}
	if attributes["country"] == "" {
		attributes["country"] = " "
	}
	if attributes["commerce"] == "" {
		attributes["commerce"] = COMMERCE
	}
	if attributes["channel"] == "" {
		attributes["channel"] = CHANNEL
	}
	if attributes["domain"] == "" {
		attributes["domain"] = DOMAIN
	}
	if attributes["capability"] == "" {
		attributes["capability"] = CAPABILITY
	}
	if attributes["mimeType"] == "" {
		attributes["mimeType"] = MIMETYPE
	}
	if attributes["sourceSystem"] == "" {
		attributes["sourceSystem"] = " "
	}
	if attributes["siteId"] == "" {
		attributes["siteId"] = " "
	}
}
