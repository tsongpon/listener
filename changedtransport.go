package main

type Transport struct {
	Object	string	`json:"object"`
	Entry []entry 	`json:"entry"`
}

type entry struct {
	Time int64 		`json:"time"`
	Change []change	`json:"changes"`
	Id string	`json:"id"`
	Uid	string `json:"uid"`
}

type change struct {
	Field	string	`json:"field"`
	Value 	string	`json:"value"`
}

