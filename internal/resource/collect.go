package resource

import "gin-chat/internal/model"

// CollectListResponse 收藏列表响应结构
type CollectListResponse struct {
	ID      int    `json:"id"`
	Type    int8   `json:"type"`
	Content string `json:"content"`
	Options string `json:"options"`
}

// CollectListResource 收藏列表转换
func CollectListResource(list []*model.CollectModel) []*CollectListResponse {
	if len(list) == 0 {
		return []*CollectListResponse{}
	}
	res := make([]*CollectListResponse, 0, len(list))
	for _, collect := range list {
		res = append(res, &CollectListResponse{
			ID:      collect.ID,
			Type:    collect.Type,
			Content: collect.Content,
			Options: collect.Options,
		})
	}
	return res
}
