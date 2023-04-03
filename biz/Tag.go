package biz

import (
	"github.com/jlu-cow-studio/common/dal/mysql"
	mysql_model "github.com/jlu-cow-studio/common/model/dao_struct/mysql"
	"gorm.io/gorm"
)

func GetTagCategoryListByMarkObject(markObject string) ([]*mysql_model.TagCategoryWithList, error) {

	list := []*mysql_model.TagCategoryWithList{}

	tx := mysql.GetDBConn().Table("tag").Preload("TagList", "mark_object = ?", markObject).
		Table("tag_category").Find(&list)

	return list, tx.Error

}

func GetTagListByItem(itemId string) ([]*mysql_model.Tag, error) {

	list := []*mysql_model.Tag{}

	tx := mysql.GetDBConn().Table("items").Where("items.id = ?", itemId).
		Joins("left join item_tag on items.id = item_tag.item_id").
		Joins("left join tag on item_tag.tag_id = tag.tag_id").
		Select("tag.*").Find(&list)

	return list, tx.Error
}

func GetTagListByUser(userId string) ([]*mysql_model.Tag, error) {
	list := []*mysql_model.Tag{}

	tx := mysql.GetDBConn().Table("user").Where("user.uid = ?", userId).
		Joins("left join user_tag on user.uid = user_tag.user_id").
		Joins("left join tag on user_tag.tag_id = tag.tag_id").
		Select("tag.*").Find(&list)

	return list, tx.Error
}

func UpdateItemTags(itemId string, tagList []string) error {

	return mysql.GetDBConn().Transaction(func(tx *gorm.DB) error {

		if err := tx.Table("item_tag").Delete(nil, "item_id = ?", itemId).Error; err != nil {
			return err
		}

		list := make([]struct {
			ItemId string  `gorm:"column:item_id"`
			TagId  string  `gorm:"column:tag_id"`
			Weight float64 `gorm:"column:weight"`
		}, len(tagList))

		for i, tagId := range tagList {
			list[i].ItemId = itemId
			list[i].TagId = tagId
			list[i].Weight = 1
		}

		if err := tx.Table("item_tag").Create(&list).Error; err != nil {
			return err
		}

		return nil
	})

}

func UpdateUserTags(userId string, tagList []string) error {
	return mysql.GetDBConn().Transaction(func(tx *gorm.DB) error {

		if err := tx.Table("user_tag").Delete(nil, "user_id = ?", userId).Error; err != nil {
			return err
		}

		list := make([]struct {
			ItemId string  `gorm:"column:user_id"`
			TagId  string  `gorm:"column:tag_id"`
			Weight float64 `gorm:"column:weight"`
		}, len(tagList))

		for i, tagId := range tagList {
			list[i].ItemId = userId
			list[i].TagId = tagId
			list[i].Weight = 1
		}

		if err := tx.Table("user_tag").Create(&list).Error; err != nil {
			return err
		}

		return nil
	})

}
