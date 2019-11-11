package user

import "model"

type User struct {
	UserId	int64	`json:"userid"`
	Name string		`json:"name"`
	Avatar string	`json:"avatar,omitempty"`
	Email string	`json:"Email,omitempty"`
	Signature string	`json:"signature,omitempty"`
	Title string		`json:"title,omitempty"`
	Group string		`json:"group,omitempty"`
	Tags []model.KLabel	`json:"tags,omitempty"`
	NotifyCount int		`json:"notifyCount,omitempty"`
	UnreadCount int		`json:"unreadCount,omitempty"`
	Country string		`json:"country,omitempty"`
	Geographic model.Geographic	`json:"geographic,omitempty"`
	Address string		`json:"address,omitempty"`
	Phone string		`json:"phone,omitempty"`
}
