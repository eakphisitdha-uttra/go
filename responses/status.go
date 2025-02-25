package responses

// 400, 404, 422, 500
type Error struct {
	Status  int         `json:"status" extensions:"x-order=0"`
	Error   string      `json:"error" extensions:"x-order=1"`
	Message string      `json:"message" extensions:"x-order=2"`
	Details interface{} `json:"details" extensions:"x-order=3"`
}

// 200, 202, 201
type Success struct {
	Status  int         `json:"status" extensions:"x-order=0"`
	Message string      `json:"message" extensions:"x-order=1"`
	Data    interface{} `json:"data" extensions:"x-order=2"`
}
