package utils

import "time"

const format = time.RFC3339Nano

func Now() time.Time {
	return time.Now()
}

func NowString() string {
	return Now().Format(format)
}

func TimeToString(t time.Time) string {
	return t.Format(format)
}

func StringToTime(v string) time.Time {
	t, err := time.Parse(format, v)
	if err != nil {
		return Now()
	}
	return t
}

func StringTimeParse(v string) (time.Time, error) {
	return time.Parse(format, v)
}

func UnmarshalBinaryTimeParse(v []byte) (time.Time, error) {
	return StringTimeParse(string(v))
}
