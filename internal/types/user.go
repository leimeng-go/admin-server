package types


type UserInfo struct {
	ID       int64 `json:"id"`
	Username string `json:"user_name"`
	Nickname string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	RoleID   int64 `json:"role_id"`
	Status   int64  `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
