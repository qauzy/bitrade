package entity

func (this *WebsiteInformation) SetId(Id int64) (result *WebsiteInformation) {
	this.Id = Id
	return this
}
func (this *WebsiteInformation) GetId() (Id int64) {
	return this.Id
}
func (this *WebsiteInformation) SetName(Name string) (result *WebsiteInformation) {
	this.Name = Name
	return this
}
func (this *WebsiteInformation) GetName() (Name string) {
	return this.Name
}
func (this *WebsiteInformation) SetLogo(Logo string) (result *WebsiteInformation) {
	this.Logo = Logo
	return this
}
func (this *WebsiteInformation) GetLogo() (Logo string) {
	return this.Logo
}
func (this *WebsiteInformation) SetAddressIcon(AddressIcon string) (result *WebsiteInformation) {
	this.AddressIcon = AddressIcon
	return this
}
func (this *WebsiteInformation) GetAddressIcon() (AddressIcon string) {
	return this.AddressIcon
}
func (this *WebsiteInformation) SetUrl(Url string) (result *WebsiteInformation) {
	this.Url = Url
	return this
}
func (this *WebsiteInformation) GetUrl() (Url string) {
	return this.Url
}
func (this *WebsiteInformation) SetKeywords(Keywords string) (result *WebsiteInformation) {
	this.Keywords = Keywords
	return this
}
func (this *WebsiteInformation) GetKeywords() (Keywords string) {
	return this.Keywords
}
func (this *WebsiteInformation) SetDescription(Description string) (result *WebsiteInformation) {
	this.Description = Description
	return this
}
func (this *WebsiteInformation) GetDescription() (Description string) {
	return this.Description
}
func (this *WebsiteInformation) SetCopyright(Copyright string) (result *WebsiteInformation) {
	this.Copyright = Copyright
	return this
}
func (this *WebsiteInformation) GetCopyright() (Copyright string) {
	return this.Copyright
}
func (this *WebsiteInformation) SetPostcode(Postcode string) (result *WebsiteInformation) {
	this.Postcode = Postcode
	return this
}
func (this *WebsiteInformation) GetPostcode() (Postcode string) {
	return this.Postcode
}
func (this *WebsiteInformation) SetContact(Contact string) (result *WebsiteInformation) {
	this.Contact = Contact
	return this
}
func (this *WebsiteInformation) GetContact() (Contact string) {
	return this.Contact
}
func (this *WebsiteInformation) SetOtherInformation(OtherInformation string) (result *WebsiteInformation) {
	this.OtherInformation = OtherInformation
	return this
}
func (this *WebsiteInformation) GetOtherInformation() (OtherInformation string) {
	return this.OtherInformation
}
func NewWebsiteInformation(id int64, name string, logo string, addressIcon string, url string, keywords string, description string, copyright string, postcode string, contact string, otherInformation string) (ret *WebsiteInformation) {
	ret = new(WebsiteInformation)
	ret.Id = id
	ret.Name = name
	ret.Logo = logo
	ret.AddressIcon = addressIcon
	ret.Url = url
	ret.Keywords = keywords
	ret.Description = description
	ret.Copyright = copyright
	ret.Postcode = postcode
	ret.Contact = contact
	ret.OtherInformation = otherInformation
	return
}

type WebsiteInformation struct {
	Id               int64  `gorm:"column:id" json:"id"`
	Name             string `gorm:"column:name" json:"name"`
	Logo             string `gorm:"column:logo" json:"logo"`
	AddressIcon      string `gorm:"column:address_icon" json:"addressIcon"`
	Url              string `gorm:"column:url" json:"url"`
	Keywords         string `gorm:"column:keywords" json:"keywords"`
	Description      string `gorm:"column:description" json:"description"`
	Copyright        string `gorm:"column:copyright" json:"copyright"`
	Postcode         string `gorm:"column:postcode" json:"postcode"`
	Contact          string `gorm:"column:contact" json:"contact"`
	OtherInformation string `gorm:"column:other_information" json:"otherInformation"`
}
