package common

type StatusCode int32

const (
	StatusCode_NO_STATUS_CODE StatusCode = 0
	StatusCode_OK             StatusCode = 200
	StatusCode_INTERNAL_ERROR StatusCode = 500
	StatusCode_BAD_REQUEST    StatusCode = 400
	StatusCode_CREATED        StatusCode = 201
)

type Response struct {
	StatusCode StatusCode
	Msg        []string
	ErrorMsg   []string
}

func GetResponse(statusCode StatusCode, msgs []string, errMsgs []string) *Response {
	return &Response{
		StatusCode: statusCode,
		ErrorMsg:   errMsgs,
		Msg:        msgs,
	}
}

func GetDefaultResponse() *Response {
	return GetResponse(StatusCode_OK, nil, nil)
}

func GetErrMsgsResponse(statusCode StatusCode, errMsgs ...string) *Response {
	return GetResponse(statusCode, nil, errMsgs)
}

func GetSuccessResponse(statusCode StatusCode, msgs ...string) *Response {
	return GetResponse(statusCode, msgs, nil)
}

func GetErrResponse(statusCode StatusCode, errs ...error) *Response {
	var errMsgs []string
	for _, err := range errs {
		errMsgs = append(errMsgs, err.Error())
	}
	return GetResponse(statusCode, nil, errMsgs)
}
