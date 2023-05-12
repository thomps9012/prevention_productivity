package methods

import (
	"context"
	"errors"
	"thomps9012/prevention_productivity/graph/model"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// queries
func FindItemNotes(item_id string, item_filter bson.D, item_type string) ([]*model.NoteDetail, error) {
	can_view_notes, err := database.Db.Collection(item_type).CountDocuments(context.TODO(), item_filter)
	if err != nil {
		return nil, err
	}
	if can_view_notes == 0 {
		return nil, errors.New("you're attempting to view notes on an item that you either didn't create, or doesn't exist")
	}
	filter := bson.D{{Key: "$match", Value: bson.D{{Key: "item_id", Value: item_id}}}}
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "author"}}}}
	unwind := bson.D{{Key: "$unwind", Value: "$author"}}
	pipeline := mongo.Pipeline{filter, user_stage, unwind}
	notes := make([]*model.NoteDetail, 0)
	cursor, err := database.Db.Collection("notes").Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}
func FindNoteDetail(filter bson.D) (*model.NoteDetail, error) {
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "author"}}}}
	unwind := bson.D{{Key: "$unwind", Value: "$author"}}
	pipeline := mongo.Pipeline{filter, user_stage, unwind}
	note := make([]*model.NoteDetail, 0)
	cursor, err := database.Db.Collection("notes").Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &note)
	if err != nil {
		return nil, err
	}
	if len(note) == 0 {
		return nil, errors.New("you're attempting to view a note that either doesn't exist, or you didn't create")
	}
	return note[0], nil
}
func FindUserNotes(user_id string) ([]*model.Note, error) {
	var notes []*model.Note
	cursor, err := database.Db.Collection("notes").Find(context.TODO(), bson.D{{Key: "user_id", Value: user_id}})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

// mutations
func CreateNote(new_note model.NewNote, note_author string, item_type string, item_filter bson.D) (*model.NoteDetail, error) {
	can_view_notes, err := database.Db.Collection(item_type).CountDocuments(context.TODO(), item_filter)
	if err != nil {
		return nil, err
	}
	if can_view_notes == 0 {
		return nil, errors.New("you're attempting to create a note for an item that you either didn't create, or doesn't exist")
	}
	collection := database.Db.Collection("notes")
	note := model.Note{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().Format("2003-01-02 15:05:05"),
		Title:     new_note.Title,
		Content:   new_note.Content,
		ItemID:    new_note.ItemID,
		UserID:    note_author,
		UpdatedAt: bson.TypeNull.String(),
	}
	res, err := collection.InsertOne(context.TODO(), note)
	if err != nil {
		return nil, err
	}
	var author_info model.UserOverview
	err = database.Db.Collection("users").FindOne(context.TODO(), bson.D{{Key: "_id", Value: note_author}}, options.FindOne().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "first_name", Value: 1}, {Key: "last_name", Value: 1}})).Decode(&author_info)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == "" {
		return nil, errors.New("failed to insert note")
	}
	return &model.NoteDetail{
		ID:        note.ID,
		ItemID:    note.ItemID,
		Author:    &author_info,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}, nil
}
func UpdateNote(update model.UpdateNote, filter bson.D) (*model.Note, error) {
	collection := database.Db.Collection("notes")
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	var note model.Note
	err := collection.FindOne(context.TODO(), filter).Decode(&note)
	if err != nil {
		return nil, errors.New("you're attempting to update a note that either doesn't exist or you didn't create")
	}
	update_args := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: update.Title},
			{Key: "content", Value: update.Content},
			{Key: "updated_at", Value: updated_at},
		}},
	}
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	var n model.Note
	err = collection.FindOneAndUpdate(context.TODO(), filter, update_args, &opts).Decode(&n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}
func DeleteNote(filter bson.D) (bool, error) {
	collection := database.Db.Collection("notes")
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return result.DeletedCount == 1, nil
}
