package models

import (
	"github.com/zing-dev/soft-version/soft"
	"time"
)

type InfoModel struct {
	StartTime string       `json:"start_time"`
	Version   soft.Version `json:"version"`
}

func GetVersion() InfoModel {
	return InfoModel{
		StartTime: time.Now().Format("2006-01-02 15:04:05"),
		Version:   soft.GetCli().Soft.Version[0],
	}
}
