package formatter

import (
	"falabella.com/boss-message-library/internal/config/log"
	"strconv"
	"strings"
	"time"
)

const TimeZoneSantiago = "America/Santiago"
const TimeZoneArgentinaBuenosAires = "America/Argentina/Buenos_Aires"
const TimeZoneColombia = "America/Bogota"
const TimeZonePeru = "America/Lima"
const TimeZoneUtc = "Etc/UTC"

const ParseFormat = "01/02/2006 15:04:05"
const ParseWithDifferentFormat = "2006-01-02 15:04:05"
const ParseFormatUtc = "2006-01-02T15:04:05.000Z"

func GetTimeZone(country string) string {
	switch country {
	case "CL":
		return TimeZoneSantiago

	case "PE":
		return TimeZonePeru

	case "CO":
		return TimeZoneColombia

	case "AR":
		return TimeZoneArgentinaBuenosAires

	default:
		return TimeZoneUtc

	}
}

var logger = log.GetLogger()

func ParseInLocation(date string, country string) string {
	if "" != date {
		location, er := time.LoadLocation(GetTimeZone(country))

		if nil != er {
			logger.Errorf("error occurred inf parse date for: %s", er)
			return ""
		}

		if strings.Contains(date, "/") {
			inLocation, err := time.ParseInLocation(ParseFormat, date, location)
			if err != nil {
				logger.Error("Error parsing date: " + date + " in location: " + GetTimeZone(country) + " error: " + err.Error())
				return ""
			}
			return inLocation.UTC().Format(ParseFormatUtc)
		}
		inLocation, err := time.ParseInLocation(ParseWithDifferentFormat, date, location)
		if err != nil {
			logger.Error("Error parsing date: " + date + " in location: " + GetTimeZone(country) + " error: " + err.Error())
			return ""
		}
		return inLocation.UTC().Format(ParseFormatUtc)

	}
	return date
}

func LocalDateTimeUTC() string {
	now := time.Now().UTC().Format(ParseFormatUtc)
	return now
}

func Timestamp() string {
	t := time.Now().UTC().UnixMilli()
	return strconv.FormatInt(t, 10)
}

func GetLocalDate(date string, country string) time.Time {

	location, er := time.LoadLocation(GetTimeZone(country))
	if er != nil {
		logger.Errorf("error load Location : %s", er)
	}
	if strings.Contains(date, "/") {
		inLocation, err := time.ParseInLocation("01/02/2006 15:04:05", date, location)
		if err != nil {
			logger.Errorf("error parsing date : %s", err)
		}
		return inLocation.UTC()
	} else if "" != date {
		//inLocation, err := time.ParseInLocation("01-02-2006 15:04:05", date, location)
		inLocation, err := time.ParseInLocation("2006-01-02 15:04:05", date, location)
		if err != nil {
			logger.Errorf("error parsing date : %s", err)
		}
		return inLocation.UTC()
	}
	return time.Time{}
}
