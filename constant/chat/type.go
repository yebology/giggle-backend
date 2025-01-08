package chat

// Chat represents different types of chat in the application.
type Chat string

// Constants for the available chat types.
var (

	Group   	Chat = "Group"    
	Personal 	Chat = "Personal" 

)

// AllowedType lists all valid chat types.
var AllowedType = []Chat{
	Group, Personal,
}
