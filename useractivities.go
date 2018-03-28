package main

type UserActivities struct {
	Total int          `json:"total"`
	Size  int          `json:"size"`
	Data  []UserChange `json:"data"`
}
