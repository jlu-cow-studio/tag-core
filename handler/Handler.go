package handler

import "github.com/jlu-cow-studio/common/dal/rpc/tag_core"

type Handler struct {
	tag_core.UnimplementedTagCoreServiceServer
}
