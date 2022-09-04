package dto

type R struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Succeed(data interface{}) R {
	return R{
		Code:    "Success",
		Message: "",
		Data:    data,
	}
}

func Reject(err error) R {
	return R{
		Code:    "Rejected",
		Message: err.Error(),
		Data:    nil,
	}
}
