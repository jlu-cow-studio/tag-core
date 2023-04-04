package handler

import (
	"context"

	"github.com/jlu-cow-studio/common/dal/rpc/base"
	"github.com/jlu-cow-studio/common/dal/rpc/tag_core"
	"github.com/jlu-cow-studio/tag-core/biz"
)

func (h *Handler) UpdateItemTags(ctx context.Context, req *tag_core.UpdateItemTagsRequest) (res *tag_core.UpdateItemTagsResponse, err error) {
	res = &tag_core.UpdateItemTagsResponse{
		Base: &base.BaseRes{
			Message: "",
			Code:    "498",
		},
	}

	if err = biz.UpdateItemTags(req.ItemId, req.TagList); err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "400"
		return res, nil
	}

	res.Base.Code = "200"
	res.Base.Message = ""

	return
}
