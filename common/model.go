package common

type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (o *Response) SetCode(code int32) {
	o.Code = code

}

func (o *Response) SetMessage(message string) {
	o.Message = message
}

func (o *Response) SetData(data interface{}) {
	o.Data = data
}

func (e Response) Clone() *Response {
	return &e
}
