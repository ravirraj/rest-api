package types

type Student struct {
	Id    int64  
	Name  string `validate:"required"`
	Email string `validate:"required"`
	Age   int    `validate:"required"`
}


type UpdateStudent struct {
	Name *string 
	Email *string
	Age *int
}