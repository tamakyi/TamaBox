// route/user/tags.go

package user

import (
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/tamakyi/TamaBox/internal/context"
	"github.com/tamakyi/TamaBox/internal/db"
	"github.com/tamakyi/TamaBox/internal/form"
)

// UpdateTags 处理标签更新表单
func UpdateTags(ctx context.Context, f form.UpdateTags) {
	// 处理标签：去重、限制长度、限制数量
	tagsStr := f.Tags
	tagsStr = strings.ReplaceAll(tagsStr, "，", ",") // 中文逗号转英文
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

		// 检查每个标签长度不超过5个字符（使用rune计数）
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

	// 更新数据库
	if err := db.Users.UpdateTags(ctx.Request().Context(), ctx.User.ID, db.UpdateTagsOptions{
		Tags:       strings.Join(words, ","),
		ShowInHome: f.ShowInHome,
	}); err != nil {
		logrus.WithContext(ctx.Request().Context()).WithError(err).Error("Failed to update tags")
		ctx.SetInternalErrorFlash()
		ctx.Redirect("/user/profile")
		return
	}

	ctx.SetSuccessFlash("更新标签成功")
	ctx.Redirect("/user/profile")
}