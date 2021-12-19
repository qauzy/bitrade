package entity

func (this *WebsiteInformation) SetId(id int64) (result *WebsiteInformation) {
	this.Id = id
	return this
}
func (this *WebsiteInformation) GetId() (id int64) {
	return this.Id
}
func (this *WebsiteInformation) SetName(name string) (result *WebsiteInformation) {
	this.Name = name
	return this
}
func (this *WebsiteInformation) GetName() (name string) {
	return this.Name
}
func (this *WebsiteInformation) SetLogo(logo string) (result *WebsiteInformation) {
	this.Logo = logo
	return this
}
func (this *WebsiteInformation) GetLogo() (logo string) {
	return this.Logo
}
func (this *WebsiteInformation) SetAddressIcon(addressIcon string) (result *WebsiteInformation) {
	this.AddressIcon = addressIcon
	return this
}
func (this *WebsiteInformation) GetAddressIcon() (addressIcon string) {
	return this.AddressIcon
}
func (this *WebsiteInformation) SetUrl(url string) (result *WebsiteInformation) {
	this.Url = url
	return this
}
func (this *WebsiteInformation) GetUrl() (url string) {
	return this.Url
}
func (this *WebsiteInformation) SetKeywords(keywords string) (result *WebsiteInformation) {
	this.Keywords = keywords
	return this
}
func (this *WebsiteInformation) GetKeywords() (keywords string) {
	return this.Keywords
}
func (this *WebsiteInformation) SetDescription(description string) (result *WebsiteInformation) {
	this.Description = description
	return this
}
func (this *WebsiteInformation) GetDescription() (description string) {
	return this.Description
}
func (this *WebsiteInformation) SetCopyright(copyright string) (result *WebsiteInformation) {
	this.Copyright = copyright
	return this
}
func (this *WebsiteInformation) GetCopyright() (copyright string) {
	return this.Copyright
}
func (this *WebsiteInformation) SetPostcode(postcode string) (result *WebsiteInformation) {
	this.Postcode = postcode
	return this
}
func (this *WebsiteInformation) GetPostcode() (postcode string) {
	return this.Postcode
}
func (this *WebsiteInformation) SetContact(contact string) (result *WebsiteInformation) {
	this.Contact = contact
	return this
}
func (this *WebsiteInformation) GetContact() (contact string) {
	return this.Contact
}
func (this *WebsiteInformation) SetOtherInformation(otherInformation string) (result *WebsiteInformation) {
	this.OtherInformation = otherInformation
	return this
}
func (this *WebsiteInformation) GetOtherInformation() (otherInformation string) {
	return this.OtherInformation
}

type WebsiteInformation struct {
	Id               int64
	Name             string
	Logo             string
	AddressIcon      string
	Url              string
	Keywords         string
	Description      string
	Copyright        string
	Postcode         string
	Contact          string
	OtherInformation string
}
