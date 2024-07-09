package models

type Request struct {
	Action string  `json:"action"`
	Body   Comment `json:"body"`
}
