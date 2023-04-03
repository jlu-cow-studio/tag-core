package handler

import (
	"context"

	"github.com/jlu-cow-studio/common/dal/rpc/base"
	"github.com/jlu-cow-studio/common/dal/rpc/tag_core"
	"github.com/jlu-cow-studio/tag-core/biz"
)

func (h *Handler) GetTagListByUser(ctx context.Context, req *tag_core.GetTagListByUserRequest) (res *tag_core.GetTagListByUserResponse, err error) {
	res = &tag_core.GetTagListByUserResponse{
		Base: &base.BaseRes{
			Message: "",
			Code:    "498",
		},
	}

	list, err := biz.GetTagListByUser(req.UserId)
	if err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "400"
		return res, nil
	}

	for _, tag := range list {
		rpcTag := &tag_core.Tag{
			Id:         tag.ID,
			Name:       tag.Name,
			Weight:     tag.Weight,
			MarkObject: tag.MarkObject,
			CategoryId: tag.CategoryID,
		}
		res.TagList = append(res.TagList, rpcTag)
	}

	res.Base.Message = ""
	res.Base.Code = "200"

	return
}
