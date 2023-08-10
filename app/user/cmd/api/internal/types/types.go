// Code generated by goctl. DO NOT EDIT.
package types

type PersonalInfo struct {
	StudentId string `json:"studentId"`
	UserName  string `json:"userName"`
	Avatar    string `json:"avatar"`
	IsBlocked int64  `json:"isBlocked"`
	RoleGrade int64  `json:"roleGrade"`
	Integral  int64  `json:"integral"`
	Licence   int64  `json:"licence"`
}

type OtherInfo struct {
	StudentId string `json:"studentId"`
	UserName  string `json:"userName"`
	Avatar    string `json:"avatar"`
}

type LoginRequest struct {
	StudentId string `json:"studentId"`
	Password  string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	IfNew bool   `json:"ifnew"`
}

type UpdateInfoRequest struct {
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}

type UpdateInfoResponse struct {
}

type SharingRequest struct {
}

type SharingResponse struct {
}

type MyInfoRequest struct {
}

type MyInfoResponse struct {
	PersonalInfo PersonalInfo `json:"personalInfo"`
}

type OtherInfoRequest struct {
	StudentID string `path:"sid"`
}

type OtherInfoResponse struct {
	OtherInfo OtherInfo `json:"otherInfo"`
}
