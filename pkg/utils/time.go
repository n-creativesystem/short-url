package utils

import "time"

func Now() time.Time {
	return time.Now()
}

func NowString() string {
	return Now().Format(time.RFC3339)
}

func TimeToString(t time.Time) string {
	return t.Format(time.RFC3339)
}

func StringToTime(v string) time.Time {
	t, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return Now()
	}
	return t
}
