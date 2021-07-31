package agentm

import (
	"container/list"
	"context"
	"gin/conn"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os/user"
	"time"
)

type Promoter struct {
	Id	primitive.ObjectID	`json:"_id" bson:"_id"`
	PromoterId	int32	`json:"promoterId,omitempty"`
}

type GameChannel struct {
	Ids	primitive.ObjectID	`json:"_id" bson:"_id"`
	Id	int32	`json:"id,omitempty"`
	DownloadURL string	`json:"downloadURL,omitempty"`
}

type ProTemCon struct {
	Id	primitive.ObjectID	`json:"_id" bson:"_id"`
	Mark	string	`json:"mark,omitempty"`
}

type RefreshQrcodeArgs struct {
	ChannelId	int32	`json:"channelId"`
	PromoterId	int32	`json:"promoterId"`
}

type InsertProTemConArgs struct {
	Id	primitive.ObjectID	`json:"_id" bson:"_id"`
	AgentId int32 `json:"agentId"`
	Note string `json:"note"`
	Mark string `json:"mark"`
	CreateTime time.Time `json:"createTime"`
}

func PromoterInfo(id int32, promoterCollection string) (Promoter, error) {
	db := conn.GetMongoDB("GAME_MAIN")
	promoter := Promoter{}
	err := db.Collection(promoterCollection).FindOne(context.TODO(),bson.D{{"promoterId",&id}}).Decode(&promoter)
	return promoter, err
}

func GameChannelInfo(id int32, gameChannelCollection string) (GameChannel, error) {
	db := conn.GetMongoDB("GAME_CONFIG")
	gameChannel := GameChannel{}
	err := db.Collection(gameChannelCollection).FindOne(context.TODO(),bson.D{{"id",&id}}).Decode(&gameChannel)
	return gameChannel, err
}

func ProTemConInfo(id int32, ProTemConCollection string) (ProTemCon, error) {
	db := conn.GetMongoDB("GAME_MAIN")
	proTemCon := ProTemCon{}
	err := db.Collection(ProTemConCollection).FindOne(context.TODO(),bson.D{{"agentId",&id}}).Decode(&proTemCon)
	return proTemCon, err
}

func InsertTemCon(insertPtcArgs InsertProTemConArgs, ProTemConCollection string) (*mongo.InsertOneResult, error) {
	db := conn.GetMongoDB("GAME_MAIN")
	res, err := db.Collection(ProTemConCollection).InsertOne(context.TODO(), bson.D{
		{"agentId",insertPtcArgs.AgentId},
		{"note",insertPtcArgs.Note},
		{"mark",insertPtcArgs.Mark},
		{"createTime",insertPtcArgs.CreateTime},
	})
	return res, err
}