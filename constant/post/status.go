package post

// Status represents the status of a post.
type Status string

// Constants for possible post statuses.
const (

	Open  Status = "Open"
	Close Status = "Closed"
	
)

// AllowedStatus lists all valid statuses.
var AllowedStatus = []Status{
	Open, Close,
}
