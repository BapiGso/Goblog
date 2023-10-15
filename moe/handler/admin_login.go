package handler

import (
	"SMOE/moe/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"net/http"
	"time"
)

// Deprecated: use JWT instead of session
func LoginGetSession(c echo.Context) error {
	//a:=c.Get("_session_store").(*sessions.Store)
	sess, err := session.Get("smoeSession", c)
	if err != nil {
		return err
	}
	if sess.Values["isLogin"] == true {
		return c.Render(http.StatusOK, "admin.template", nil)
	}
	return c.Render(http.StatusOK, "login.template", nil)
}

func LoginGet(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.Render(http.StatusOK, "login.template", nil)
	}
	if user.Valid { // 校验token
		slog.Info("someone login")
		return c.Render(http.StatusOK, "admin.template", nil)
	}
	return echo.ErrUnauthorized
}

func LoginPost(c echo.Context) error {
	qpu := database.NewQPU()
	defer qpu.Free()
	req := &struct {
		Name     string `form:"user" validate:"required,min=1,max=200"`
		Pwd      string `form:"pwd" validate:"required,min=8,max=200"`
		Illsions string `form:"illsions" `
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := qpu.UserWithName(req.Name); err != nil {
		return err
	}
	//计算提交表单的密码与盐 scrypt和数据库中密码是否一致
	if err := bcrypt.CompareHashAndPassword([]byte(qpu.UserInfo.Password), []byte(req.Pwd)); err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), //过期日期设置7天
		})
		t, err := token.SignedString(SigningKey)
		if err != nil {
			return err
		}
		c.SetCookie(&http.Cookie{
			Name:     "smoe_token",
			Value:    t,
			HttpOnly: true,
		})
		return c.Redirect(302, "/admin")
	}
	//TODO 发邮件提醒和防爆破
	return echo.ErrUnauthorized
}
