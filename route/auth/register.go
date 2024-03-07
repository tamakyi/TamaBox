// Copyright 2022 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/flamego/recaptcha"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/tamakyi/TamaBox/internal/conf"
	"github.com/tamakyi/TamaBox/internal/context"
	"github.com/tamakyi/TamaBox/internal/db"
	"github.com/tamakyi/TamaBox/internal/form"
)

func Register(ctx context.Context) {
	ctx.Success("auth/register")
}

func RegisterAction(ctx context.Context, f form.Register, recaptcha recaptcha.RecaptchaV3) {
	if ctx.HasError() {
		ctx.Success("auth/register")
		return
	}

	// Check recaptcha code.
	resp, err := recaptcha.Verify(f.Recaptcha, ctx.Request().Request.RemoteAddr)
	if err != nil {
		logrus.WithContext(ctx.Request().Context()).WithError(err).Error("Failed to check recaptcha")
		ctx.SetInternalErrorFlash()
		ctx.Redirect("/register")
		return
	}
	if !resp.Success {
		ctx.SetErrorFlash("验证码错误")
		ctx.Redirect("/register")
		return
	}

	if err := db.Users.Create(ctx.Request().Context(), db.CreateUserOptions{
		Name:                 f.Name,
		Password:             f.Password,
		Email:                f.Email,
		Avatar:               conf.Upload.DefaultAvatarURL,
		Domain:               f.Domain,
		Background:           conf.Upload.DefaultBackground,
		Intro:                "问你想问的",
		Qrcodebackcolor:      "#000000",
		Qrcodecolor:          "#696969",
		Dotscale:             "0.5",
		BackgroundImage:      conf.Upload.DefaultBackgroundImageURL,
		Backgroundimagealpha: "0.5",
		Qrcodepdpcolor:       "#696969",
	}); err != nil {
		switch {
		case errors.Is(err, db.ErrUserNotExists),
			errors.Is(err, db.ErrBadCredential),
			errors.Is(err, db.ErrDuplicateEmail),
			errors.Is(err, db.ErrDuplicateDomain):
			ctx.SetError(errors.Cause(err))

		default:
			logrus.WithContext(ctx.Request().Context()).WithError(err).Error("Failed to create new user")
			ctx.SetInternalError()
		}

		ctx.Success("auth/register")
		return
	}

	ctx.SetSuccessFlash("注册成功，欢迎来到 狼的提问箱！")
	ctx.Redirect("/login")
}
