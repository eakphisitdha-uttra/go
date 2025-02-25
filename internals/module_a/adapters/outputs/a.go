package outputs

type GetOutput struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Field string `json:"field"`
	//
	// add your data
	//
}

type ValidateOutput struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
