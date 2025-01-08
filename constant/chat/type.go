package chat

type Chat string

var (

	Group 		Chat = "Group"
	Personal 	Chat = "Personal"

)

var AllowedType = []Chat{
	Group, Personal,
}