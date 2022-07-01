package utils

import (
	"github.com/sirupsen/logrus"
	"time"
)

func CurrentTime() string {
	now := time.Now().UTC()
	return now.Format("2006-01-02 15:04:05")
}

func TimeDuration(startT, endT string) uint {
	sT, _ := time.Parse("2006-01-02 15:04:05", startT)
	eT, _ := time.Parse("2006-01-02 15:04:05", endT)
	durT := eT.Sub(sT).Seconds()
	logrus.Infof("from: %v to: %v  dur:%v", sT, eT, durT)
	return uint(durT)
}
