package param_bind

type ProductAdd struct {
	Name string `form:"name" json:"name" validate:"required,NameValid"`
}
