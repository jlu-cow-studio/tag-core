package handler

import (
	"context"

	"github.com/jlu-cow-studio/common/dal/rpc/base"
	"github.com/jlu-cow-studio/common/dal/rpc/tag_core"
	"github.com/jlu-cow-studio/tag-core/biz"
)

func (h *Handler) UpdateUserTags(ctx context.Context, req *tag_core.UpdateUserTagsRequest) (res *tag_core.UpdateUserTagsResponse, err error) {
	res = &tag_core.UpdateUserTagsResponse{
		Base: &base.BaseRes{
			Message: "",
			Code:    "498",
		},
	}

	if err = biz.UpdateUserTags(req.UserId, req.TagList); err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "400"
		return res, nil
	}

	res.Base.Code = "200"
	res.Base.Message = ""

	return
}
