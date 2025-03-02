package inputs

type AddInput struct {
	Name string `json:"name" binding:"required"`
	//
	// add your data
	//
}

type UpdateInput struct {
	ID int
	//
	// add your data
	//
}

type DeleteInput struct {
	ID int
	//
	// add your data
}
