package utc

import "time"

func GetInfoUTC() (date string, datetime string, epoch int64) {
	ahora := time.Now().UTC()
	date = ahora.Format("2006-01-02")
	datetime = ahora.Format("2006-01-02 15:04:05")
	epoch = ahora.UnixNano() / int64(time.Millisecond)
	return
}
