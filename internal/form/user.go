// Copyright 2022 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package form

type UpdateProfile struct {
	Name                 string `valid:"required;maxlen:20" label:"昵称"`
	OldPassword          string `label:"旧密码"`
	NewPassword          string `valid:"maxlen:30" label:"新密码"`
	Intro                string `valid:"required;maxlen:100" label:"介绍"`
    Usernamecolor        string `valid:"required;minlen:7;maxlen:7" label:"用户名颜色"`
    Introcolor           string `valid:"required;minlen:7;maxlen:7" label:"个人说明颜色"`
	Qrcodebackcolor      string `valid:"required;minlen:7;maxlen:7" label:"二维码背景色"`
	Qrcodecolor          string `valid:"required;minlen:7;maxlen:7" label:"二维码前景色"`
	Dotscale             string `valid:"required" label:"二维码点大小"`
	Backgroundimagealpha string `valid:"required" label:"二维码背景透明度"`
	Qrcodepdpcolor       string `valid:"required;minlen:7;maxlen:7" label:"二维码探测图形颜色"`
	NotifyEmail          string `label:"开启邮箱通知"`
}

type UpdateHarassment struct {
	RegisterOnly string `label:"仅允许注册用户"`
	BlockWords string `label:"屏蔽词"`
}
