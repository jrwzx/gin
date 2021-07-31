package user

import (
	"errors"
	"fmt"
	help "gin/common/help"
	user "gin/models/users"
	"github.com/gin-gonic/gin"
	"net/http"
)
const GameUserCollection = "game_user"
var (
	errNotExist        = errors.New("用户不存在")
	errUpdationFailed  = errors.New("更新失败")
)

func GetAllUser(c *gin.Context) {
	gameUsers, err := user.GetAllUser(GameUserCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &gameUsers})
}

func GetUser(c *gin.Context) {
	userId := c.Param("userId")
	userId32 := help.StringToInt32(userId)
	userInfo, err := user.UserInfo(userId32, GameUserCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &userInfo})
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("userId")
	userId32 := help.StringToInt32(userId)
	_, err := user.UserInfo(userId32, GameUserCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
		return
	}
	args := user.UpdateArgs{UserId32: userId32, Fields: "万千万", GameUserCollection: GameUserCollection}
	result, err := user.UpdateUser(args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errUpdationFailed.Error()})
		return
	}
	if result.MatchedCount != 0 {
		c.JSON(http.StatusOK, gin.H{"status":"success", "message":fmt.Sprintf("更新成功%v个",result.MatchedCount )})
	}else{
		c.JSON(http.StatusOK, gin.H{"status":"failed", "message":"更新失败"})
	}
}
//func CreateUser(c *gin.Context) {
//	// Get DB from Mongo Config
//	db := conn.GetMongoDB()
//	user := user.User{}
//	err := c.Bind(&user)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
//		return
//	}
//	user.ID = bson.NewObjectId()
//	user.CreatedAt = time.Now()
//	user.UpdatedAt = time.Now()
//	err = db.C(GameUserCollection).Insert(user)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &user})
//}



//func DeleteUser(c *gin.Context) {
//	// Get DB from Mongo Config
//	db := conn.GetMongoDB()
//	var id ObjectId = ObjectIdHex(c.Param("id")) // Get Param
//	err := db.C(GameUserCollection).Remove(M{"_id": &id})
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errDeletionFailed.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User deleted successfully"})
//}