package agent

import (
	"encoding/json"
	"errors"
	"fmt"
	user "gin/models/users"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"strings"

	//"fmt"
	agent "gin/models/agentm"
	"github.com/gin-gonic/gin"
	"github.com/xxtea/xxtea-go/xxtea"
	"net/http"
	"time"
)

const PromoterCollection = "promoter"
const GameChannelCollection = "game_channel"
const ProTemConCollection = "promoter_template_configs"
const Char = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ChannelEncryptKey = ""

var (
	errParams = errors.New("参数错误")
	errNotExist = errors.New("数据不存在")
	errFailParse = errors.New("解析错误")
	errTemConFail = errors.New("生成模板配置失败")
)

func RefreshQrcode(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	params := string(buf[0:n])
	bytesData := []byte(params)
	var argsOj agent.RefreshQrcodeArgs
	var mark string
	if json.Unmarshal(bytesData,&argsOj) == nil {
		_, errPromoter := agent.PromoterInfo(argsOj.PromoterId, PromoterCollection)
		if errPromoter != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
			return
		}
		gameChannel, errGameChannel := agent.GameChannelInfo(argsOj.ChannelId, GameChannelCollection)
		if errGameChannel != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
			return
		}

		_, errProTemCon := agent.ProTemConInfo(argsOj.PromoterId, ProTemConCollection) //proTemConData
		if errProTemCon != nil {
			for{
				info := make(map[string]interface{})
				info["promoterId"] = argsOj.PromoterId
				info["channelId"] = argsOj.ChannelId
				info["rand"] = time.Now()
				infoJson, errInfo := json.Marshal(info)
				if errInfo != nil {
					c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errFailParse.Error()})
					return
				}else{
					mark = string(xxtea.Encrypt([]byte(infoJson),[]byte(ChannelEncryptKey)))
					if strings.Contains(mark,"+") || strings.Contains(mark,"/") {
						break
					}
				}
			}
			randFirst := RandomStr(rand.Intn(2)+3)
			randSecond := RandomStr(rand.Intn(2)+3)
			promotionUrl := gameChannel.DownloadURL + string(randFirst) + ".html?" + string(randSecond) + "=" + mark + "&channelId=" + string(argsOj.ChannelId)
			timeString := time.Now().Format("2006-01-02 15:04:05")
			note := "生成于" + timeString

			insertPtcArgs := agent.InsertProTemConArgs{
					AgentId:argsOj.PromoterId,
					Note:note,
					Mark:mark,
					CreateTime:time.Now(),
			}
			_, err := agent.InsertTemCon(insertPtcArgs,ProTemConCollection)
			if err != nil{
				c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errTemConFail.Error()})
				return
			}
		}else{

		}




	}else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errParams.Error()})
		return
	}
}

func RandomStr(length int) ([]byte){
	bytesChar := []byte(Char)
	var str []byte
	for i := length; i>0; i-- {
		str = append(str,bytesChar[rand.Intn(len(bytesChar))])
	}
	return str
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