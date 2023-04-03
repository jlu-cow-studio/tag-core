package biz

import (
	"fmt"
	"testing"

	"github.com/jlu-cow-studio/common/dal/mysql"
	"github.com/sanity-io/litter"
)

func TestGetTagCategoryListByMarkObject(t *testing.T) {
	mysql.Init()

	list, err := GetTagCategoryListByMarkObject("interest")
	litter.Dump(list)
	fmt.Println(len(list), err)
}

func TestGetTagListByItem(t *testing.T) {
	mysql.Init()

	list, err := GetTagListByItem("1")
	litter.Dump(list)
	fmt.Println(err)
}

func TestGetTagListByUser(t *testing.T) {
	mysql.Init()

	list, err := GetTagListByUser("322")
	litter.Dump(list)
	fmt.Println(err)
}

func TestUpdateItemTag(t *testing.T) {
	mysql.Init()

	err := UpdateItemTags("1", []string{"62", "30", "63", "19", "28", "1", "10"})

	fmt.Println(err)
}

func TestUpdateUserTag(t *testing.T) {
	mysql.Init()

	err := UpdateUserTags("322", []string{"27", "49", "18"})

	fmt.Println(err)
}
