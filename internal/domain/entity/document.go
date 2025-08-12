package entity

type Document struct {
	ID          int    `gorm:"column:ID"`
	Name        string `gorm:"column:Name"`
	IsClaim     int    `gorm:"column:isClaim"`
	IsRefund    int    `gorm:"column:isRefund"`
	CreatedBy   string `gorm:"column:Created_By"`
	CreatedAt   string `gorm:"column:Created_At"`
	UpdatedBy   string `gorm:"column:Updated_By"`
	UpdatedAt   string `gorm:"column:Updated_At"`
	Usage       int    `gorm:"column:Usage"`
	TotalRecord int    `gorm:"column:TotalRecord"`
}
