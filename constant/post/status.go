package post

type Status string

const (

	Open Status = "Open"
	Close Status = "Closed"

)

var AllowedStatus = []Status{
	Open, Close,
}