package resource

import "gin-chat/internal/model"

// ApplyListResponse 申请列表响应结构
type ApplyListResponse struct {
	User   *model.User `json:"user"`
	Status int8        `json:"status"`
}

// ApplyListResource 申请列表转化
func ApplyListResource(applys []*model.ApplyModel, users []*model.UserModel) []*ApplyListResponse {
	um := usersToMap(users)
	list := make([]*ApplyListResponse, len(applys))
	for _, apply := range applys {
		if user, ok := um[apply.UserID]; ok {
			list = append(list, &ApplyListResponse{
				User:   user,
				Status: apply.Status,
			})
		}
	}
	return list
}
