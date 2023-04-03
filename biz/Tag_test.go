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
