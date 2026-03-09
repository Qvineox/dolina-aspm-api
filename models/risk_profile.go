package models

type RiskProfile struct {
	ID   uint32 `gorm:"column:id; primaryKey; autoIncrement"`
	SUID string `gorm:"column:suid; uniqueIndex; not null"`

	Name        string `gorm:"column:name; uniqueIndex; not null"`
	Description string `gorm:"column:description"`
}
