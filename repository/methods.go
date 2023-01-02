package repo

import (
	"context"
	"log"

	"github.com/dvnvln/deallscrud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) connect(ctx context.Context) (*mongo.Database, error) {
	err := r.client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.Database(r.dbName), nil
}

func (r *Repository) Get(ctx context.Context) ([]model.User, error) {
	db, err := r.connect(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("getting user data")
	cursor, err := db.Collection(model.UserCollection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var result []model.User = nil

	err = cursor.All(ctx, &result)
	return result, err
}

func (r *Repository) GetByUserID(ctx context.Context, userID string) ([]model.User, error) {
	db, err := r.connect(ctx)
	if err != nil {
		return nil, err
	}
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	log.Println("getting user data: ", objID)
	cursor, err := db.Collection(model.UserCollection).Find(ctx, bson.M{model.BsonObjID: objID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var result []model.User = nil
	err = cursor.All(ctx, &result)
	return result, err
}

func (r *Repository) Add(ctx context.Context, user model.User) error {
	db, err := r.connect(ctx)
	if err != nil {
		return err
	}
	_, err = db.Collection(model.UserCollection).InsertOne(ctx, user)
	if err != nil {
		return err
	}
	log.Println("data added")
	return nil
}

func (r *Repository) Update(ctx context.Context, userID string, user model.User) error {
	db, err := r.connect(ctx)
	if err != nil {
		return err
	}
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	log.Println("getting user data: ", objID)
	user.ID = objID
	_, err = db.Collection(model.UserCollection).UpdateOne(ctx, bson.M{model.BsonObjID: objID}, bson.M{"$set": user})
	return err
}

func (r *Repository) Delete(ctx context.Context, userID string) error {
	db, err := r.connect(ctx)
	if err != nil {
		return err
	}
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = db.Collection(model.UserCollection).DeleteOne(ctx, bson.M{model.BsonObjID: objID})
	return err
}

func (r *Repository) Login(ctx context.Context, user model.UserReqBody) (model.User, error) {
	db, err := r.connect(ctx)
	if err != nil {
		return model.User{}, err
	}
	log.Print("repo here..")
	uname := user.Username
	var result model.User = model.User{}
	err = db.Collection(model.UserCollection).FindOne(ctx, bson.M{model.BsonUname: uname}).Decode(&result)
	log.Print("repo done here..")
	return result, err
}
