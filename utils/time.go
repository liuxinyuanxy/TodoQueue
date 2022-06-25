package utils

import "time"

func CurrentTime() string {
	now := time.Now().UTC()
	return now.Format("2006-01-02 15:04:05")
}

func TimeDuration(startT, endT string) uint {
	sT, _ := time.Parse("", startT)
	eT, _ := time.Parse("", endT)
	durT := eT.Sub(sT)
	return uint(durT)
}
