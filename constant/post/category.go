package post

type Category string

const (

	UiUxDesign Category = "UI/UX Design"
	WebDev Category = "Web Development"
	AppDev Category = "App Development"
	Web3 Category = "Web3"
	GraphicDesign Category = "Graphic Design"
	Animation Category = "Animation"
	Seo Category = "SEO"
	Copywriting Category = "Copywriting"
	DigitalMarketing Category = "Digital Marketing"
	PhotoVideoEditing Category = "Photo/Video Editing"

)

var AllowedCategories = []Category{
	UiUxDesign, WebDev, AppDev, Web3,
	GraphicDesign, Animation, Seo,
	Copywriting, DigitalMarketing, PhotoVideoEditing,
}