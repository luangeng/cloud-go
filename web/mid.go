package web

// import (
// 	"errors"
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// func Cors(c *gin.Context) {
// 	method := c.Request.Method
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Headers", "Content-Type,X-CSRF-Token, lg-token,lg-client")
// 	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
// 	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
// 	c.Header("Access-Control-Allow-Credentials", "true")
// 	if method == "OPTIONS" {
// 		c.AbortWithStatus(http.StatusNoContent)
// 	}
// 	c.Next()
// }

// func queryToken(c *gin.Context) (dao.Session, error) {
// 	token := c.GetHeader(com.SessionKey)
// 	if len(strings.TrimSpace(token)) != 15 {
// 		token, _ = c.Cookie(com.SessionKey)
// 	}
// 	if len(strings.TrimSpace(token)) != 15 {
// 		return dao.Session{}, errors.New("")
// 	}
// 	data, err := dao.QuerySession(token)
// 	if err != nil {
// 		return dao.Session{}, err
// 	}
// 	return data, nil
// }

// func TokenFilter(c *gin.Context) {
// 	session, err := queryToken(c)
// 	if err != nil {
// 		c.JSON(200, dao.Error("未登录或登录过期"))
// 		c.Abort()
// 		return
// 	}
// 	c.Set("token", session.ID)
// 	c.Set("id", session.UserId)
// 	c.Set("name", session.Name)
// 	c.Set("role", session.Role)
// 	c.Next()
// }

// func AdminFilter(c *gin.Context) {
// 	role := c.GetString("role")
// 	if role != "admin" {
// 		c.JSON(200, dao.Error("无权访问"))
// 		c.Abort()
// 		return
// 	}
// 	c.Next()
// }

// func ApiFilter(c *gin.Context) {
// 	token := c.GetHeader("token")
// 	if token == "" {
// 		c.JSON(200, dao.Error("no token"))
// 		return
// 	}
// 	userKey := tool.AesDecrypt(token, tool.Config.TokenPrivateKey)
// 	var user dao.User
// 	err := dao.GetDB().Where("key = ?", userKey).First(&user).Error
// 	if err != nil {
// 		c.JSON(200, dao.Error("invalid token"))
// 		return
// 	}
// 	c.Set("id", user.ID)
// }

// var NodejsCallerLimitCh = make(chan struct{}, 10)

// func limit(c *gin.Context) {
// 	NodejsCallerLimitCh <- struct{}{}
// 	defer func() { <-NodejsCallerLimitCh }()
// }
