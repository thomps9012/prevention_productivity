package methods

import (
	"context"
	"errors"
	"thomps9012/prevention_productivity/graph/model"
	database "thomps9012/prevention_productivity/internal/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindItemNotes(item_id string) ([]*model.NoteDetail, error) {
	return nil, errors.New("method unimplemented")
}
func FindNoteDetail(note_id string) (*model.NoteDetail, error) {
	return nil, errors.New("method unimplemented")
}
func FindUserNotes(user_id string) ([]*model.Note, error) {
	return nil, errors.New("method unimplemented")
}

func CreateNote(new_note model.NewNote, note_author string) (*model.NoteDetail, error) {
	collection := database.Db.Collection("notes")
	note := model.Note{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().Format("2003-01-02 15:05:05"),
		UpdatedAt: bson.TypeNull.String(),
	}
	res, err := collection.InsertOne(context.TODO(), note)
	if err != nil {
		return nil, err
	}
	var author_info model.UserOverview
	err = database.Db.Collection("users").FindOne(context.TODO(), bson.D{{"_id", note_author}}, options.FindOne().SetProjection(bson.D{{"_id", 1}, {"first_name", 1}, {"last_name", 1}})).Decode(&author_info)
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
	err := collection.FindOneAndUpdate(context.TODO(), filter, update_args, &opts).Decode(&n)
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
