package error_wrapper

type ErrorException struct {
	HttpCode int         `json:"-"`
	Errno    string      `json:"errno"`
	DataType string      `json:"dataType"`
	ErrorMsg string      `json:"error"`
	Data     interface{} `json:"data"`
	Request  string      `json:"-"`
}

func NewErrorException(httpCode int, errno string, dataType string, errorMsg string, data interface{}) *ErrorException {
	return &ErrorException{HttpCode: httpCode, Errno: errno, DataType: dataType, ErrorMsg: errorMsg, Data: data}
}

func (e *ErrorException) Error() string {
	return e.ErrorMsg
}
