package answer

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sh4ll0t/db"
)

func DeleteAnswer(c *gin.Context) {
	session := sessions.Default(c)
	if auth, ok := session.Get("authenticated").(bool); !ok || !auth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	id := c.PostForm("id")
	var answer db.Answer
	if err := db.DB.First(&answer, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if session.Get("username") != answer.Respondent {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "只有本作者才可以删除！"})
		return
	}

	db.DB.Delete(&db.Answer{}, "id = ?", id)
	db.DB.Delete(&answer)
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
