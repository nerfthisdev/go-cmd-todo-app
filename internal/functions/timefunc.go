package functions

import "time"

func GetTimeFromZone(zone string) (string, error) {
	location, err := time.LoadLocation(zone)

	if err != nil {
		return "", err
	}

	timeNow := time.Now().In(location)
	return timeNow.Format(time.RFC1123), nil
}
