// Package user provides types/client for making requests to /me.
package user

// User represents the currently authenticated user.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Avatar   string `json:"avatar"`
	IsPublic bool   `json:"isPublic"`
}

// Operation represents operations a user can perform and their quotas.
type Operation struct {
	Name    string `json:"name"`
	Limit   int    `json:"limit"`
	Usage   int    `json:"usage"`
	Overage int    `json:"overage"`
}

type authenticatedUserWrapper struct {
	User       User        `json:"user"`
	Operations []Operation `json:"operations"`
}
