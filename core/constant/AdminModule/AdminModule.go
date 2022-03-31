package AdminModule

var (
	CMS       = AdminModule{"CMS", 0}
	COMMON    = AdminModule{"COMMON", 1}
	EXCHANGE  = AdminModule{"EXCHANGE", 2}
	FINANCE   = AdminModule{"FINANCE", 3}
	MEMBER    = AdminModule{"MEMBER", 4}
	OTC       = AdminModule{"OTC", 5}
	SYSTEM    = AdminModule{"SYSTEM", 6}
	PROMOTION = AdminModule{"PROMOTION", 7}
	INDEX     = AdminModule{"INDEX", 8}
	IEO       = AdminModule{"IEO", 9}
	GIFT      = AdminModule{"GIFT", 10}
	MARGIN    = AdminModule{"MARGIN", 11}
)

func (this *AdminModule) Ordinal() (result int) {
	return this.ordinal
}

type AdminModule struct {
	Title   string `gorm:"column:title" json:"title"`
	ordinal int
}
