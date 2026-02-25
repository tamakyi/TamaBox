// route/user/explore.go

package user

import (
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/tamakyi/TamaBox/internal/context"
	"github.com/tamakyi/TamaBox/internal/db"
)

// exploreUser 用于模板展示，包含分割后的标签切片
type exploreUser struct {
	*db.User
	TagList []string
}

// Explore 显示所有启用了 show_in_home 的用户，并支持按标签搜索
func Explore(ctx context.Context) {
	keyword := ctx.Query("q")
	var users []*db.User
	var err error

	if keyword != "" {
		users, err = db.Users.GetByTag(ctx.Request().Context(), keyword)
	} else {
		users, err = db.Users.GetShowInHome(ctx.Request().Context())
	}
	if err != nil {
		logrus.WithContext(ctx.Request().Context()).WithError(err).Error("Failed to get users for explore")
		ctx.SetInternalError()
		ctx.Success("user/explore")
		return
	}

	// 构建带有分割标签的列表
	exploreUsers := make([]*exploreUser, 0, len(users))
	for _, u := range users {
		var tags []string
		if u.Tags != "" {
			tags = strings.Split(u.Tags, ",")
		}
		exploreUsers = append(exploreUsers, &exploreUser{
			User:    u,
			TagList: tags,
		})
	}

	ctx.Data["Users"] = exploreUsers
	ctx.Data["Keyword"] = keyword
	ctx.SetTitle("探索用户标签 - 狼的提问箱")
	ctx.Success("user/explore")
}