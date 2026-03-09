package models

type Vendor struct {
	ID   uint32 `gorm:"column:id; primaryKey; autoIncrement"`
	Name string `gorm:"column:name; uniqueIndex:uidx_vendor_name; not null"`
	URL  string `gorm:"column:url"`

	VendorScorings []VendorScoring `gorm:"foreignKey:VendorName; references:Name; constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
}

func (v Vendor) TableName() string {
	return "advisory_vendors"
}
