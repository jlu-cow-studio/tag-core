package handler

import (
	"context"
	"fmt"
	"testing"

	"github.com/jlu-cow-studio/common/dal/mysql"
	"github.com/jlu-cow-studio/common/dal/redis"
	"github.com/jlu-cow-studio/common/dal/rpc/base"
	"github.com/jlu-cow-studio/common/dal/rpc/tag_core"
	"github.com/jlu-cow-studio/common/discovery"
	"github.com/sanity-io/litter"
)

func TestGetTagListByScene(t *testing.T) {

	discovery.Init()
	redis.Init()
	mysql.Init()

	req := &tag_core.GetTagListBySceneRequest{
		Base:  &base.BaseReq{},
		Scene: "interest",
	}

	res, err := new(Handler).GetTagListByScene(context.Background(), req)
	litter.Dump(res)
	fmt.Println(err)
}
