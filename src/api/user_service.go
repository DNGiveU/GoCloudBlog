package api

import (
	"constants"
	"github.com/gin-gonic/gin"
	"model"
	"model/user"
	"net/http"
)

var visitor user.User = user.User{
	UserId: -1,
	Name: "游客",
}

/**
 * 获取当前用户
 */
func GetCurrentUser(c *gin.Context) {
	cookie, err := c.Cookie(constants.COOKIE_TOKEN)
	if err != nil {
		// 游客
		c.JSON(http.StatusOK, visitor)
		return
	}
	println("cookie: " + cookie)

	user := user.User{
		UserId: 1,
		Name: "NGiveU",
		Avatar: "https://avatars3.githubusercontent.com/u/17162085?s=40&v=4",
		Signature: "海纳百川，有容乃大",
		Title: "Go 开发",
		Group: "工程开发师",
		Tags: []model.KLabel{
			{
				Key: "0",
				Label: "很有想法的",
			},
			{
				Key: "1",
				Label: "专注设计与开发",
			},
		},
		NotifyCount: 12,
		UnreadCount: 11,
		Country: "China",
		Geographic: model.Geographic{
			Province: model.KLabel{
				Key: "330000",
				Label: "四川省",
			},
			City: model.KLabel{
				Key:   "330100",
				Label: "成都市",
			},
		},
		Address: "成都市",
		Phone: "666",
	}

	c.JSON(http.StatusOK, user)
}
