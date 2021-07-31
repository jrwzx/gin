package users

import (
	"context"
	"gin/conn"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"log"
)
type RegInfo struct {
	OsType    int32        	`json:"osType"`
}

type GameUser struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	UserId    int32        	`json:"userId,omitempty"`
	TrueName  string        `json:"trueName,omitempty"`
	RegInfo	  RegInfo		`json:"regInfo,omitempty"`
}

type UpdateArgs struct {
	UserId32 int32
	Fields string
	GameUserCollection string
}

func UserInfo(id int32, userCollection string) (GameUser, error) {
	db := conn.GetMongoDB("GAME_MAIN")
	gameUser := GameUser{}
	err := db.Collection(userCollection).FindOne(context.TODO(),bson.D{{"userId",&id}}).Decode(&gameUser)
	return gameUser, err
}

func GetAllUser(userCollection string) ([] GameUser, error) {
	var GameUsers [] GameUser
	db := conn.GetMongoDB("GAME_MAIN")
	cursor, err1 := db.Collection(userCollection).Find(context.TODO(), bson.D{{}})
	if err1 != nil {
		log.Fatal(err1)
	}
	err2 := cursor.All(context.TODO(),&GameUsers)
	return GameUsers, err2
}

func UpdateUser(updateArgs UpdateArgs) (*mongo.UpdateResult, error) {
	db := conn.GetMongoDB("GAME_MAIN")
	filter := bson.D{{"userId", updateArgs.UserId32}}
	update := bson.D{{"$set", bson.D{{"trueName", &updateArgs.Fields}}}}
	result, err := db.Collection(updateArgs.GameUserCollection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}