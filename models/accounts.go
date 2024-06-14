package models

type LoginModel struct {
	Permanent     bool       `json:"_permanent"`
	Authenticated bool       `json:"authenticated"`
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	Role          int        `json:"role"`
	Perms         int        `json:"perms"`
	Template      string     `json:"template"`
	Flashes       [][]string `json:"_flashes"`
}

type AllUsersDataModel map[string]struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Role         int    `json:"role"`
	Permission   int    `json:"permission"`
	TemplateName string `json:"template_name"`
}
