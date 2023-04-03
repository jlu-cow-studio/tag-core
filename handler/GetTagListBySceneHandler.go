package handler

import (
	"context"

	"github.com/jlu-cow-studio/common/dal/rpc/base"
	"github.com/jlu-cow-studio/common/dal/rpc/tag_core"
	mysql_model "github.com/jlu-cow-studio/common/model/dao_struct/mysql"
	"github.com/jlu-cow-studio/tag-core/biz"
)

func (h *Handler) GetTagListByScene(ctx context.Context, req *tag_core.GetTagListBySceneRequest) (res *tag_core.GetTagListBySceneResponse, err error) {
	res = &tag_core.GetTagListBySceneResponse{
		Base: &base.BaseRes{
			Message: "",
			Code:    "498",
		},
	}

	if _, ok := mysql_model.MarkObjectMap[req.Scene]; !ok {
		res.Base.Message = "unknown scene!"
		res.Base.Code = "400"
		return res, nil
	}

	list, err := biz.GetTagCategoryListByMarkObject(req.Scene)
	if err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "401"
		return res, nil
	}

	res.List = make([]*tag_core.TagCategoryWithList, len(list))

	for i, tag := range list {
		tagpb := &tag_core.TagCategoryWithList{
			Category: &tag_core.TagCategory{
				Id:       int32(tag.ID),
				Name:     tag.Name,
				ParentId: int32(tag.ParentID),
				Level:    int32(tag.Level),
			},
			TagList: make([]*tag_core.Tag, len(tag.TagList)),
		}

		for j, subTag := range tag.TagList {
			tagpb.TagList[j] = &tag_core.Tag{
				Id:         subTag.ID,
				Name:       subTag.Name,
				Weight:     subTag.Weight,
				MarkObject: subTag.MarkObject,
				CategoryId: subTag.CategoryID,
			}
		}

		res.List[i] = tagpb
	}

	res.Base.Message = ""
	res.Base.Code = "200"

	return
}
