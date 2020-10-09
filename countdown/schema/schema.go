package schema

import (
	"context"
	"fmt"
	"log"

	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
var col = client.Database("count").Collection("timer")

type User struct{
	ID primitive.ObjectID `bson :"_id,omitempty"`
	Name string `bson:"name,omitempty"`
	Email string `bson:"email,omitempty"`
	Password string `+bson:"password,omitempty"`
	Events []Event `bson:"event,omitempty"`
}

type Event struct{
	Name string
	Date int
}
func Construct(name,email,password string)User{
	u:=User{}
	u.Name=name
	u.Email=email
	u.Password= password
	return u
}
func (u *User)Insert(){
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	client.Connect(ctx)
	col.InsertOne(ctx,bson.M{"_id":u.ID,"name":u.Name,"email":u.Email,"password":u.Password,"event":u.Events})
}

func (u*User)Get()bson.M{
	var user bson.M
	data := make([]bson.M, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	cur, err := col.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("not")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var epi bson.M
		if err = cur.Decode(&epi); err != nil {
			log.Fatal(err)
		}
		data = append(data, epi)
	}
	for _, elem := range data {
		if elem["name"] == u.Name {
			if elem["password"]==u.Password{
				user = elem 
			}else{
				fmt.Println("password doesnt math")
			}				
		}else{
			fmt.Println("you are not present in our  database")
		}
	}
	
	fmt.Println(user)
	return user
	// return user["name"].(string),user["email"].(string),user["password"].(string),user["event"].([]Event)
}
func (u *User)Addevent(e Event){
	u.Events =append(u.Events,e)
}

func Update(u1  User){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	result,err := col.UpdateOne(ctx,bson.M{"_id":u1.ID},bson.M{"$set":bson.M{"event":u1.Events}})
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(result.ModifiedCount)
}
func GetAll()[]bson.M{
	data := make([]bson.M, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	cur, err := col.Find(context.TODO(), bson.M{"val": 8})
	if err != nil {
		fmt.Println("not")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var epi bson.M
		if err = cur.Decode(&epi); err != nil {
			log.Fatal(err)
		}
		data = append(data, epi)
	}
	return data
}
func insertone(val int) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	col.InsertOne(ctx, bson.M{"val": val})
}
func get(val int) {
	data := make([]bson.M, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	cur, err := col.Find(context.TODO(), bson.M{"val": 8})
	if err != nil {
		fmt.Println("not")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var epi bson.M
		if err = cur.Decode(&epi); err != nil {
			log.Fatal(err)
		}
		data = append(data, epi)
	}
	for _, elem := range data {
		if elem["val"] == val {
			fmt.Println(elem)
		}
}

	fmt.Println(data)
}
