package models

import (
	soft "github.com/zing-dev/soft-version/src"
	"time"
)

type InfoModel struct {
	StartTime string       `json:"start_time"`
	Version   soft.Version `json:"version"`
}

var Info = InfoModel{
	StartTime: time.Now().Format("2006-01-02 15:04:05"),
	Version:   soft.NewSoft().Version[0],
}
