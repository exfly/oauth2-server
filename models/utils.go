package models

// OauthUserToUser convert OauthUser to User
func OauthUserToUser(ou *OauthUser) *User {
	ret := User{
		ID:        ou.ID,
		CreatedAt: ou.CreatedAt,
		UpdatedAt: ou.UpdatedAt,
		Username:  ou.Username,
	}
	if ou.RoleID.Valid {
		ret.Role = ou.RoleID.String
	}
	return &ret
}
