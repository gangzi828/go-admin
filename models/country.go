package models

import (
	orm "go-admin/database"
)

type Country struct {
	PkId        int    `gorm:"primary_key;AUTO_INCREMENT" json:"pkId"` //自增ID
	CountryName string `gorm:"type:varchar(128);" json:"countryName"`  //国家名称
	CountryCode string `gorm:"type:varchar(128);" json:"countryCode"`  //国家代码
	StateId     string `gorm:"type:int(11);" json:"stateId"`           //所属大洲
	Remark      string `gorm:"type:varchar(255);" json:"remark"`       //描述
	CreateBy    string `gorm:"type:varchar(128);" json:"createBy"`
	UpdateBy    string `gorm:"type:varchar(128);" json:"updateBy"`
	BaseModel
}

func (Country) TableName() string {
	return "geography_country"
}

func (e *Country) GetPage(pageSize int, pageIndex int) ([]Country, int, error) {
	var doc []Country

	table := orm.Eloquent.Select("*").Table(e.TableName())

	if e.CountryName != "" {
		table = table.Where("country_name = ?", e.CountryName)
	}
	if e.CountryCode != "" {
		table = table.Where("country_code = ?", e.CountryCode)
	}

	var count int

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}
