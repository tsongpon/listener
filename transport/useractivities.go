package transport

import "github.com/tsongpon/listener/data"

type UserActivities struct {
	Total int               `json:"total"`
	Size  int               `json:"size"`
	Data  []data.UserChange `json:"data"`
}
