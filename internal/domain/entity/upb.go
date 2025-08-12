package entity

import "time"

type Upb struct {
	ID                int       `gorm:"column:ID"`
	Code              string    `gorm:"column:Code"`
	Name              string    `gorm:"column:Name"`
	Email             string    `gorm:"column:Email"`
	Phone             string    `gorm:"column:Phone"`
	Fax               string    `gorm:"column:Fax"`
	Lat               float64   `gorm:"column:Lat"`
	Long              float64   `gorm:"column:Long"`
	Address           string    `gorm:"column:Address"`
	Province          string    `gorm:"column:Province"`
	City              string    `gorm:"column:City"`
	District          string    `gorm:"column:District"`
	Village           string    `gorm:"column:Village"`
	Active            int       `gorm:"column:Active"`
	Officer1          string    `gorm:"column:Officer1"`
	Officer2          string    `gorm:"column:Officer2"`
	Description       string    `gorm:"column:Description"`
	CustomerID        int       `gorm:"column:Customer_ID"`
	Customer          string    `gorm:"column:Customer"`
	Alias             string    `gorm:"column:Alias"`
	EmailDueLetter    string    `gorm:"column:EmailDueLetter"`
	EmailDueLetterCC  string    `gorm:"column:EmailDueLetterCC"`
	EmailDueLetterCC2 string    `gorm:"column:EmailDueLetterCC2"`
	CreatedBy         string    `gorm:"column:Created_By"`
	CreatedAt         time.Time `gorm:"column:Created_At"`
	UpdatedBy         string    `gorm:"column:Updated_By"`
	UpdatedAt         time.Time `gorm:"column:Updated_At"`
	DeletedBy         string    `gorm:"column:Deleted_By"`
	DeletedAt         time.Time `gorm:"column:Deleted_At"`
	DeleteStatus      int       `gorm:"column:DeleteStatus"`
}
