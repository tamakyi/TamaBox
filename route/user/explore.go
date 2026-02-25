// Copyright 2026 tamakyi. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"strings"
    "strconv"
	"github.com/sirupsen/logrus"

	"github.com/tamakyi/TamaBox/internal/context"
	"github.com/tamakyi/TamaBox/internal/db"
)

var tagColors = []string{
	"uk-label-success",
	"uk-label-warning",
	"uk-label-danger",
	"uk-label-primary",
	"uk-label-secondary",
}

type exploreUser struct {
	*db.User
	TagList []string
	TagColors []string
}

func Explore(ctx context.Context) {
    keyword := ctx.Query("q")
    pageStr := ctx.Query("page")
    page := 1
    if pageStr != "" {
        if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
            page = p
        }
    }
    pageSize := 20

    var users []*db.User
    var total int64
    var err error

    if keyword != "" {
        users, total, err = db.Users.GetByTag(ctx.Request().Context(), keyword, page, pageSize)
    } else {
        users, total, err = db.Users.GetShowInHome(ctx.Request().Context(), page, pageSize)
    }
    if err != nil {
        logrus.WithContext(ctx.Request().Context()).WithError(err).Error("Failed to get users for explore")
        ctx.ServerError()
        return
    }

    // 计算总页数
    totalPages := int(total) / pageSize
    if int(total)%pageSize != 0 {
        totalPages++
    }

    exploreUsers := make([]*exploreUser, 0, len(users))
    for _, u := range users {
        var tags []string
        if u.Tags != "" {
            tags = strings.Split(u.Tags, ",")
        }
        tagColorsList := make([]string, len(tags))
        for i := range tags {
            tagColorsList[i] = tagColors[i%len(tagColors)]
        }
        exploreUsers = append(exploreUsers, &exploreUser{
            User:      u,
            TagList:   tags,
            TagColors: tagColorsList,
        })
    }

    ctx.Data["Users"] = exploreUsers
    ctx.Data["Keyword"] = keyword
    ctx.Data["Page"] = page
    ctx.Data["TotalPages"] = totalPages
    ctx.Data["PrevPage"] = page - 1
    ctx.Data["NextPage"] = page + 1
    ctx.SetTitle("探索用户标签 - 狼的提问箱")
    ctx.Success("user/explore")
}