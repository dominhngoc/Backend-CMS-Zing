package response

func DataResponse(code string, message interface{}, data interface{}, detailerror interface{}, success string) (Response) {
	v := Response{}
	v.Code = code
	v.Message = message
	v.Data = data
	v.DetailErrorCode = detailerror
	v.Success = success
	return v
}
