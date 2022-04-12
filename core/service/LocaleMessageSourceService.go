package service

func (this *LocaleMessageSourceService) GetMessage(code string) (result string) {
	return this.GetMessageV2(code, nil)
}
func (this *LocaleMessageSourceService) GetMessageV2(code string, args []interface{}) (result string) {
	return this.GetMessageV3(code, args, "")
}
func (this *LocaleMessageSourceService) GetMessageV3(code string, args []interface{}, defaultMessage string) (result string) {
	var locale = LocaleContextHolder.GetLocale()
	return this.MessageSource.GetMessage(code, args, defaultMessage, locale)
}
func NewLocaleMessageSourceService(messageSource *context.MessageSource) (ret *LocaleMessageSourceService) {
	ret = new(LocaleMessageSourceService)
	ret.MessageSource = messageSource
	return
}

type LocaleMessageSourceService struct {
	MessageSource *context.MessageSource
}
