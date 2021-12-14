package entity

func (this *WebsiteInformation) SetId(id int64) {
	this.Id = id
}
func (this *WebsiteInformation) GetId() (id int64) {
	return this.Id
}
func (this *WebsiteInformation) SetName(name string) {
	this.Name = name
}
func (this *WebsiteInformation) GetName() (name string) {
	return this.Name
}
func (this *WebsiteInformation) SetLogo(logo string) {
	this.Logo = logo
}
func (this *WebsiteInformation) GetLogo() (logo string) {
	return this.Logo
}
func (this *WebsiteInformation) SetAddressIcon(addressIcon string) {
	this.AddressIcon = addressIcon
}
func (this *WebsiteInformation) GetAddressIcon() (addressIcon string) {
	return this.AddressIcon
}
func (this *WebsiteInformation) SetUrl(url string) {
	this.Url = url
}
func (this *WebsiteInformation) GetUrl() (url string) {
	return this.Url
}
func (this *WebsiteInformation) SetKeywords(keywords string) {
	this.Keywords = keywords
}
func (this *WebsiteInformation) GetKeywords() (keywords string) {
	return this.Keywords
}
func (this *WebsiteInformation) SetDescription(description string) {
	this.Description = description
}
func (this *WebsiteInformation) GetDescription() (description string) {
	return this.Description
}
func (this *WebsiteInformation) SetCopyright(copyright string) {
	this.Copyright = copyright
}
func (this *WebsiteInformation) GetCopyright() (copyright string) {
	return this.Copyright
}
func (this *WebsiteInformation) SetPostcode(postcode string) {
	this.Postcode = postcode
}
func (this *WebsiteInformation) GetPostcode() (postcode string) {
	return this.Postcode
}
func (this *WebsiteInformation) SetContact(contact string) {
	this.Contact = contact
}
func (this *WebsiteInformation) GetContact() (contact string) {
	return this.Contact
}
func (this *WebsiteInformation) SetOtherInformation(otherInformation string) {
	this.OtherInformation = otherInformation
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
