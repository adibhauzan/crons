package entity

type ApplicationSetting struct {
	ID                int    `gorm:"column:ID"`
	CompanyCode       string `gorm:"column:Company_Code"`
	CompanyName       string `gorm:"column:Company_Name"`
	CompanyAddress    string `gorm:"column:Company_Address"`
	CompanyPrintName  string `gorm:"column:Company_PrintName"`
	Phone1            string `gorm:"column:Phone1"`
	Phone2            string `gorm:"column:Phone2"`
	Email             string `gorm:"column:Email"`
	Website           string `gorm:"column:Website"`
	Logo              string `gorm:"column:Logo"`
	Sertificate1      string `gorm:"column:Sertificate1"`
	Sertificate2      string `gorm:"column:Sertificate2"`
	Sertificate3      string `gorm:"column:Sertificate3"`
	Sertificate4      string `gorm:"column:Sertificate4"`
	Sertificate5      string `gorm:"column:Sertificate5"`
	TimeOutDuration   int    `gorm:"column:TimeOut_Duration"`
	COAVAT            string `gorm:"column:COA_VAT"`
	COACommissionRev  string `gorm:"column:COA_CommissionRev"`
	COAOtherAttribute string `gorm:"column:COA_OtherAttribute"`
	COAUncollectDebt  string `gorm:"column:COA_UncollectDebt"`
	MaxAge            int    `gorm:"column:MaxAge"`
}
