package model

type GetUserByIdReqUri struct {
	ID int64 `uri:"id"`
}

type AddNewUserReqBody struct {
	Name string `json:"name"`
}

type DeleteUserReqUri struct {
	ID int64 `uri:"id"`
}
