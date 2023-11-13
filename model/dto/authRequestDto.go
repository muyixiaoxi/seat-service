package dto

type OldRolePolicy struct {
	OldRoleName string `json:"old_role_name" binding:"required"`
	OldUrl      string `json:"old_url" binding:"required"`
	OldMethod   string `json:"old_method" binding:"required"`
	RoleName    string `json:"role_name" binding:"required"`
	Url         string `json:"url" binding:"required"`
	Method      string `json:"method" binding:"required"`
}
