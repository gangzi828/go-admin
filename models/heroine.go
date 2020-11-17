package models

import (
	orm "go-admin/database"
	"time"
)

const (
	DateFormat string = "2006-01-02"
)

type Date time.Time

func (t *Date) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+DateFormat+`"`, string(data), time.Local)
	*t = Date(now)
	return
}
func (t *Date) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(DateFormat)+2)
	b = append(b, '"')
	b = time.Time(*t).AppendFormat(b, DateFormat)
	b = append(b, '"')
	return b, nil
}

type Heroine struct {
	PkId          int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name          string `gorm:"type:varchar(128);" json:"name"`
	Era           int    `gorm:"type:tinyint(2);" json:"era"`
	Birth         Date   `gorm:"type:date;" json:"birth"`
	Remark        string `gorm:"type:text;" json:"remark"`
	Encyclopedia  string `gorm:"type:varchar(255);" json:"encyclopedia"`
	Avatar        string `gorm:"type:varchar(128);" json:"avatar"`
	Height        int    `gorm:"type:int(11);" json:"height"`
	Weight        int    `gorm:"type:int(11);" json:"weight"`
	Constellation string `gorm:"type:varchar(128);" json:"constellation"`
	BaseModel
}

func (Heroine) TableName() string {
	return "heroine"
}

func (e *Heroine) Create() (Heroine, error) {
	var doc Heroine
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

func (e *Heroine) GetPage(pageSize int, pageIndex int) ([]Heroine, int, error) {
	var doc []Heroine

	table := orm.Eloquent.Select("*").Table(e.TableName())

	if e.Name != "" {
		table = table.Where("name = ?", e.Name)
	}

	var count int

	if err := table.Order("created_at desc").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}
