// Copyright 2026 tamakyi. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/tamakyi/TamaBox/internal/context"
	"github.com/tamakyi/TamaBox/internal/db"
	"github.com/tamakyi/TamaBox/internal/form"
)

func UpdateTags(ctx context.Context, f form.UpdateTags) {

	tagsStr := f.Tags
	tagsStr = strings.ReplaceAll(tagsStr, "，", ",")
	tagsStr = strings.TrimSpace(tagsStr)

	words := make([]string, 0)
	wordSet := make(map[string]struct{})
	for _, tag := range strings.Split(tagsStr, ",") {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}
		if _, ok := wordSet[tag]; ok {
			continue
		}
		wordSet[tag] = struct{}{}

		if len([]rune(tag)) > 5 {
			ctx.SetErrorFlash("每个标签不能超过 5 个字符：" + tag)
			ctx.Redirect("/user/profile")
			return
		}
		words = append(words, tag)
	}

	if len(words) > 5 {
		ctx.SetErrorFlash("标签不能超过 5 个")
		ctx.Redirect("/user/profile")
		return
	}

	showInHome := false
	if f.ShowInHome == "on" || f.ShowInHome == "1" || f.ShowInHome == "true" {
		showInHome = true
	}

	if err := db.Users.UpdateTags(ctx.Request().Context(), ctx.User.ID, db.UpdateTagsOptions{
		Tags:       strings.Join(words, ","),
		ShowInHome: showInHome,
	}); err != nil {
		logrus.WithContext(ctx.Request().Context()).WithError(err).Error("Failed to update tags")
		ctx.SetInternalErrorFlash()
		ctx.Redirect("/user/profile")
		return
	}

	ctx.SetSuccessFlash("更新标签成功")
	ctx.Redirect("/user/profile")
}