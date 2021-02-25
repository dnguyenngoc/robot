package exceptions

type BadRequest struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type InternalServerError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"status internal server error"`
}

type Success struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"status success"`
}

type NotFound struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"status not found"`
}