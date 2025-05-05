package dtos

type FindAllUsersQuery struct {
	UserNameLike string `json:"username_like"`
  Active			 bool   `json:"active"`
	
}
	