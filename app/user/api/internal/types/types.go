// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoRequest struct {
	Uid int64 `form:"uid"`
}

type UserInfoResponse struct {
	Uid   int64  `json:"uid"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}
