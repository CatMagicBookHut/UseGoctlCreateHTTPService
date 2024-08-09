package model

import (
	"fmt"
)

// Cat对象
type Cat struct {
	Uid  uint   `gorm:"primaryKey" json:"uid"`
	Name string `gorm:"varchar(100);not null" json:"name"`
	Age  int    `gorm:"int(10);not null" json:"age"`
}

// 新增猫猫信息
func CreateCat(cat Cat) int {
	err := db.Create(&cat).Error
	if err != nil {
		return 500
	}
	return 200
}

// 查询猫猫信息
func QueryCat(uid uint) (Cat, int) {
	var cat Cat
	err := db.Where("uid=?", uid).Find(&cat).Error
	if err != nil || cat.Name == "" {
		fmt.Println("查询失败")
		return cat, 500
	}
	return cat, 200
}

// 修改猫猫信息
func EditCat(cat Cat) int {
	var catMaps = make(map[string]interface{})
	catMaps["uid"] = cat.Uid
	catMaps["name"] = cat.Name
	catMaps["age"] = cat.Age
	var c Cat
	err = db.Model(&c).Where("uid=?", cat.Uid).Updates(catMaps).Error
	if err != nil {
		return 500
	}
	return 200
}

// 删除猫猫信息
func DeleteCat(uid uint) int {
	var cat Cat
	_, e := QueryCat(uid)
	if e != 200 {
		return e
	}
	err := db.Where("uid=?", uid).Unscoped().Delete(&cat).Error
	if err != nil {
		return 500
	}
	return 200
}
