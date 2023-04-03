package biz

import (
	"github.com/jlu-cow-studio/common/dal/mysql"
	mysql_model "github.com/jlu-cow-studio/common/model/dao_struct/mysql"
)

func GetTagCategoryListByMarkObject(markObject string) ([]*mysql_model.TagCategoryWithList, error) {

	list := []*mysql_model.TagCategoryWithList{}

	tx := mysql.GetDBConn().Table("tag").Preload("TagList", "mark_object = ?", markObject).
		Table("tag_category").Find(&list)

	return list, tx.Error

}
