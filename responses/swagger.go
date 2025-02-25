package responses

type SuccessOKSwagger struct {
	Status  int      `json:"code" example:"200" extensions:"x-order=0"`
	Message string   `json:"message" example:"Get successfully" extensions:"x-order=1"`
	Data    struct{} `json:"data" extensions:"x-order=2"`
}

type SuccessCreatedSwagger struct {
	Status  int      `json:"code" example:"201" extensions:"x-order=0"`
	Message string   `json:"message" example:"created successfully" extensions:"x-order=1"`
	Data    struct{} `json:"data" extensions:"x-order=2"`
}

type SuccessAcceptedSwagger struct {
	Status  int      `json:"code" example:"202" extensions:"x-order=0"`
	Message string   `json:"message" example:"Accepted successfully" extensions:"x-order=1"`
	Data    struct{} `json:"data" extensions:"x-order=2"`
}

type ErrorInternalServerErrorSwagger struct {
	Status  int      `json:"code" example:"500" extensions:"x-order=0"`
	Error   string   `json:"error" example:"Internal Server Error" extensions:"x-order=1"`
	Message string   `json:"message" example:"internal server error" extensions:"x-order=2"`
	Details struct{} `json:"details" extensions:"x-order=3"`
}

type ErrorUnauthorizedSwagger struct {
	Status  int      `json:"code" example:"401" extensions:"x-order=0"`
	Error   string   `json:"error" example:"unauthorized to access the resource" extensions:"x-order=1"`
	Message string   `json:"message" example:"unauthorized to access the resource" extensions:"x-order=2"`
	Details struct{} `json:"details" extensions:"x-order=3"`
}

type ErrorForbiddenSwagger struct {
	Status  int      `json:"code" example:"403" extensions:"x-order=0"`
	Error   string   `json:"error" example:"forbidden to access the resource" extensions:"x-order=1"`
	Message string   `json:"message" example:"forbidden to access the resource" extensions:"x-order=2"`
	Details struct{} `json:"details" extensions:"x-order=3"`
}

type ErrorConflictedSwagger struct {
	Status  int      `json:"code" example:"409" extensions:"x-order=0"`
	Error   string   `json:"error" example:"Conflict to access the resource" extensions:"x-order=1"`
	Message string   `json:"message" example:"Conflict to access the resource" extensions:"x-order=2"`
	Details struct{} `json:"details" extensions:"x-order=3"`
}

type ErrorBadRequestedSwagger struct {
	Status  int      `json:"code" example:"400" extensions:"x-order=0"`
	Error   string   `json:"error" example:"Bad Request" extensions:"x-order=1"`
	Message string   `json:"message" example:"bad request error message" extensions:"x-order=2"`
	Details struct{} `json:"details" extensions:"x-order=3"`
}

type ErrorNotFoundSwagger struct {
	Status  int      `json:"code" example:"404" extensions:"x-order=0"`
	Error   string   `json:"error" example:"Not Found" extensions:"x-order=1"`
	Message string   `json:"message" example:"data was not found" extensions:"x-order=2"`
	Details struct{} `json:"details" extensions:"x-order=3"`
}

type ErrorValidatedSwagger struct {
	Status  int      `json:"code" example:"422" extensions:"x-order=0"`
	Error   string   `json:"error" example:"Unprocessable Entity" extensions:"x-order=1"`
	Message string   `json:"message" example:"Validation failed" extensions:"x-order=2"`
	Details struct{} `json:"details" extensions:"x-order=3"`
}
