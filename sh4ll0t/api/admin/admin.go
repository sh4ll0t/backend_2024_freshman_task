package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"hduhelp_text/api/user"
	"hduhelp_text/db"
	"net/http"
)

func Admin(c *gin.Context) {
	session := sessions.Default(c)
	if auth, ok := session.Get("authenticated").(bool); !ok || !auth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}
	if session.Get("username") != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "管理员才可以访问！"})
		return
	}
	var questions []user.Question
	if err := db.DB.Preload("Answers").Find(&questions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, questions)

}
