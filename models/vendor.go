package models

type Vendor struct {
	ID uint32 `gorm:"column:id; primaryKey; autoIncrement"`

	Name string `gorm:"column:name; uniqueIndex; not null"`
	URL  string `gorm:"column:url"`
}
