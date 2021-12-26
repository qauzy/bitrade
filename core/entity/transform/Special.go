package transform

func (this *Special) SetContext(context []E) (result *Special) {
	this.Context = context
	return this
}
func (this *Special) GetContext() (context []E) {
	return this.Context
}

type Special struct {
	Context []E
}
