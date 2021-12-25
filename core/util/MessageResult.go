package util

func NewMessageResultV2(code int, msg string) (this *MessageResult) {
	this = new(MessageResult)
	this.Code = code
	this.Message = msg
	return
}
func NewMessageResultV4(code int, msg string, object interface{}, total int64) (this *MessageResult) {
	this = new(MessageResult)
	this.Code = code
	this.Message = msg
	this.Total = total
	this.Data = object
	return
}
func NewMessageResultV3(code int, msg string, object interface{}) (this *MessageResult) {
	this = new(MessageResult)
	this.Code = code
	this.Message = msg
	this.Data = object
	return
}
func NewMessageResult() (this *MessageResult) {
	this = new(MessageResult)
	// TODO Auto-generated constructor stub
	return
}
func Success() (result *MessageResult) {
	return NewMessageResultV2(0, "SUCCESS")
}
func SuccessV1(msg string) (result *MessageResult) {
	return NewMessageResultV2(0, msg)
}
func SuccessV2(msg string, data interface{}) (result *MessageResult) {
	return NewMessageResultV3(0, msg, data)
}
func SuccessDataAndTotal(object interface{}, total int64) (result *MessageResult) {
	return NewMessageResultV4(0, "SUCCESS", object, total)
}
func ErrorV2(code int, msg string) (result *MessageResult) {
	return NewMessageResultV2(code, msg)
}
func Error(msg string) (result *MessageResult) {
	return NewMessageResultV2(500, msg)
}
func (this *MessageResult) GetCode() (result int) {
	return this.Code
}
func (this *MessageResult) SetCode(code int) {
	this.Code = code
}
func (this *MessageResult) GetMessage() (result string) {
	return this.Message
}
func (this *MessageResult) SetMessage(message string) {
	this.Message = message
}
func (this *MessageResult) GetTotal() (result int64) {
	return this.Total
}
func (this *MessageResult) SetTotal(total int64) {
	this.Total = total
}

func (this *MessageResult) GetData() (result interface {
}) {
	return this.Data
}
func (this *MessageResult) SetData(data interface {
}) {
	this.Data = data
}
func GetSuccessInstance(message string, data interface{}) (result *MessageResult) {
	result = SuccessV1(message)
	result.SetData(data)
	return result
}

type MessageResult struct {
	Data    interface{}
	Code    int
	Message string
	Total   int64
}
