package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	generated1 "thomps9012/prevention_productivity/graph/generated"
	"thomps9012/prevention_productivity/graph/model"
	"thomps9012/prevention_productivity/internal/auth"
	"thomps9012/prevention_productivity/internal/methods"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *mutationResolver) CreateUser(ctx context.Context, newUser model.NewUser) (*model.LoginRes, error) {
	res, err := methods.CreateUser(newUser)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) Login(ctx context.Context, login model.LoginInput) (*model.LoginRes, error) {
	res, err := methods.LoginUser(login)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, updateUser model.UpdateUser, id string) (*model.UserUpdateRes, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: updateUser.ID}}
	}
	res, err := methods.UpdateUser(updateUser, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.UserUpdateRes, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}}
	}
	res, err := methods.DeleteUser(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) CreateGrant(ctx context.Context, newGrant model.NewGrant) (*model.GrantDetail, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return nil, admin_err
	}
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	res, err := methods.CreateGrant(newGrant, *user_id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateGrant(ctx context.Context, updateGrant model.UpdateGrant) (*model.Grant, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return nil, admin_err
	}
	res, err := methods.UpdateGrant(updateGrant)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteGrant(ctx context.Context, id string) (bool, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return false, admin_err
	}
	res, err := methods.DeleteGrant(id)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) CreateContact(ctx context.Context, newContact model.NewContact) (*model.ContactDetail, error) {
	contact_creator, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	res, err := methods.CreateContact(newContact, *contact_creator)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateContact(ctx context.Context, updateContact model.UpdateContact) (*model.Contact, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	var filter bson.D
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: updateContact.ID}, {Key: "created_by", Value: *user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: updateContact.ID}}
	}
	res, err := methods.UpdateContact(updateContact, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteContact(ctx context.Context, id string) (bool, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return false, err
	}
	admin_err := auth.ForAdmin(ctx)
	var filter bson.D
	if admin_err == nil {
		filter = bson.D{{Key: "_id", Value: id}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}, {Key: "created_by", Value: *user_id}}
	}
	res, err := methods.DeleteContact(filter)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) CreateNote(ctx context.Context, newNote model.NewNote) (*model.NoteDetail, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	res, err := methods.CreateNote(newNote, *user_id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateNote(ctx context.Context, updateNote model.UpdateNote) (*model.Note, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	is_admin := auth.ForAdmin(ctx)
	var filter bson.D
	if is_admin == nil {
		filter = bson.D{{Key: "_id", Value: updateNote.ID}}
	} else {
		filter = bson.D{{Key: "_id", Value: updateNote.ID}, {Key: "user_id", Value: *user_id}}
	}
	res, err := methods.UpdateNote(updateNote, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteNote(ctx context.Context, id string) (bool, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return false, err
	}
	is_admin := auth.ForAdmin(ctx)
	var filter bson.D
	if is_admin == nil {
		filter = bson.D{{Key: "_id", Value: id}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}, {Key: "user_id", Value: *user_id}}
	}
	res, err := methods.DeleteNote(filter)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) CreateLog(ctx context.Context, newLog model.NewLog) (*model.LogRes, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	res, err := methods.CreateNewLog(newLog, *user_id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateLog(ctx context.Context, updateLog model.UpdateLog) (*model.Log, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: updateLog.ID}, {Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: updateLog.ID}}
	}
	res, err := methods.UpdateLog(updateLog, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteLog(ctx context.Context, id string) (bool, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return false, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: id}, {Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}}
	}
	res, err := methods.DeleteLog(filter)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) ApproveLog(ctx context.Context, id string) (bool, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return false, admin_err
	}
	res, err := methods.ApproveLog(id)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) RejectLog(ctx context.Context, id string) (bool, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return false, admin_err
	}
	res, err := methods.RejectLog(id)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) CreateEvent(ctx context.Context, newEvent model.NewEvent) (*model.EventRes, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	res, err := methods.CreateEvent(newEvent, *user_id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, updateEvent model.UpdateEvent) (*model.Event, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: updateEvent.ID}, {Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: updateEvent.ID}}
	}
	res, err := methods.UpdateEvent(updateEvent, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteEvent(ctx context.Context, id string) (bool, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return false, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: id}, {Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}}
	}
	res, err := methods.DeleteEvent(filter)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) ApproveEvent(ctx context.Context, id string) (bool, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return false, admin_err
	}
	res, err := methods.ApproveEvent(id)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) RejectEvent(ctx context.Context, id string) (bool, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return false, admin_err
	}
	res, err := methods.RejectEvent(id)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) CreateEventSummary(ctx context.Context, newEventSummary model.NewEventSummary) (*model.EventSummaryRes, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	res, err := methods.CreateEventSummary(newEventSummary, *user_id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateEventSummary(ctx context.Context, updateEventSummary model.UpdateEventSummary) (*model.EventSummary, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	var filter bson.D
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: updateEventSummary.ID}, {Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: updateEventSummary.ID}}
	}
	res, err := methods.UpdateEventSummary(updateEventSummary, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteEventSummary(ctx context.Context, id string) (bool, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return false, err
	}
	var filter bson.D
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: id}, {Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}}
	}
	res, err := methods.DeleteEventSummary(filter)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) ApproveEventSummary(ctx context.Context, id string) (bool, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return false, admin_err
	}
	res, err := methods.ApproveEventSummary(id)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) RejectEventSummary(ctx context.Context, id string) (bool, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return false, admin_err
	}
	res, err := methods.RejectEventSummary(id)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) CreateSchoolReportPlan(ctx context.Context, newSchoolReportPlan model.NewSchoolReportPlan) (*model.SchoolReportPlanRes, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	res, err := methods.CreateSchoolReportPlan(newSchoolReportPlan, *user_id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateSchoolReportPlan(ctx context.Context, updateSchoolReportPlan model.UpdateSchoolReportPlan) (*model.SchoolReportPlan, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	var filter bson.D
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: updateSchoolReportPlan.ID}, {Key: "user_id", Value: *user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: updateSchoolReportPlan.ID}}
	}
	res, err := methods.UpdateSchoolReportPlan(updateSchoolReportPlan, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteSchoolReportPlan(ctx context.Context, id string) (bool, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return false, err
	}
	var filter bson.D
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: id}, {Key: "user_id", Value: *user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}}
	}
	res, err := methods.DeleteSchoolReportPlan(filter)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) ApproveSchoolReportPlan(ctx context.Context, id string) (bool, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return false, admin_err
	}
	res, err := methods.ApproveSchoolReportPlan(id)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) RejectSchoolReportPlan(ctx context.Context, id string) (bool, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return false, admin_err
	}
	res, err := methods.RejectSchoolReportPlan(id)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) CreateSchoolReportDebrief(ctx context.Context, newSchoolReportDebrief model.NewSchoolReportDebrief) (*model.SchoolReportDebriefRes, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	res, err := methods.CreateSchoolReportDebrief(newSchoolReportDebrief, *user_id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateSchoolReportDebrief(ctx context.Context, updateSchoolReportDebrief model.UpdateSchoolReportDebrief) (*model.SchoolReportDebrief, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	var filter bson.D
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: updateSchoolReportDebrief.ID}, {Key: "user_id", Value: *user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: updateSchoolReportDebrief.ID}}
	}
	res, err := methods.UpdateSchoolReportDebrief(updateSchoolReportDebrief, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteSchoolReportDebrief(ctx context.Context, id string) (bool, error) {
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return false, err
	}
	var filter bson.D
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: id}, {Key: "user_id", Value: *user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}}
	}
	res, err := methods.DeleteDebrief(filter)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) ApproveSchoolReportDebrief(ctx context.Context, id string) (bool, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return false, admin_err
	}
	res, err := methods.ApproveDebrief(id)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *mutationResolver) RejectSchoolReportDebrief(ctx context.Context, id string) (bool, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return false, admin_err
	}
	res, err := methods.RejectDebrief(id)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]model.UserResult, error) {
	admin_err := auth.ForAdmin(ctx)
	var res []model.UserResult
	var err error
	if admin_err != nil {
		res, err = methods.GetUserOverviews()
	} else {
		res, err = methods.GetUsers()
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return nil, errors.New("method unimplemented")
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return nil, errors.New("method unimplemented")
}

func (r *queryResolver) ItemNotes(ctx context.Context, itemID string, itemType string) ([]*model.NoteDetail, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: itemID}, {Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: itemID}}
	}
	res, err := methods.FindItemNotes(itemID, filter, itemType)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Note(ctx context.Context, id string) (*model.NoteDetail, error) {
	res, err := methods.FindNoteDetail(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) UserNotes(ctx context.Context, userID string) ([]*model.Note, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return nil, admin_err
	}
	res, err := methods.FindUserNotes(userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Log(ctx context.Context, id string) (*model.LogWithNotes, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: id}, {Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}}
	}
	res, err := methods.FindLogDetail(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) AllLogs(ctx context.Context) ([]*model.LogOverview, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{}}
	}
	res, err := methods.FindAllLogs(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) UserLogs(ctx context.Context, userID string) ([]*model.LogOverview, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return nil, admin_err
	}
	res, err := methods.FindUserLogs(userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Event(ctx context.Context, id string) (*model.EventWithNotes, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: id}, {Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}}
	}
	res, err := methods.FindEventDetails(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Events(ctx context.Context) ([]*model.EventOverview, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{}}
	}
	res, err := methods.FindEvents(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) UserEvents(ctx context.Context, userID string) ([]*model.EventOverview, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return nil, admin_err
	}
	res, err := methods.FindUserEvents(userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) EventSummary(ctx context.Context, id string) (*model.EventSummaryWithNotes, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: id}, {Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}}
	}
	res, err := methods.EventSummaryDetail(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) EventSummaries(ctx context.Context) ([]*model.EventSummaryOverview, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{}}
	}
	res, err := methods.FindEventSummaries(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) UserEventSummaries(ctx context.Context, userID string) ([]*model.EventSummaryOverview, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return nil, admin_err
	}
	res, err := methods.FindUserEventSummaries(userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) SchoolReportPlan(ctx context.Context, id string) (*model.SchoolReportPlanWithNotes, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: id}, {Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}}
	}
	res, err := methods.FindSchoolReportPlanDetail(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) SchoolReportPlans(ctx context.Context) ([]*model.SchoolReportPlanOverview, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{}}
	}
	res, err := methods.FindSchoolReportPlans(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) UserSchoolReportPlans(ctx context.Context, userID string) ([]*model.SchoolReportPlanOverview, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return nil, admin_err
	}
	res, err := methods.FindUserSchoolReportPlans(userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) SchoolReportDebrief(ctx context.Context, id string) (*model.SchoolReportDebriefWithNotes, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: id}, {Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}}
	}
	res, err := methods.FindSchoolReportDebriefDetail(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) SchoolReportDebriefs(ctx context.Context) ([]*model.SchoolReportDebriefOverview, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "user_id", Value: user_id}}
	} else {
		filter = bson.D{{}}
	}
	res, err := methods.FindSchoolReportDebriefs(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) UserSchoolReportDebriefs(ctx context.Context, userID string) ([]*model.SchoolReportDebriefOverview, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return nil, admin_err
	}
	res, err := methods.FindUserSchoolReportDebriefs(userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Grant(ctx context.Context, id string) (*model.GrantDetail, error) {
	res, err := methods.FindGrantDetail(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Grants(ctx context.Context) ([]*model.GrantOverview, error) {
	res, err := methods.FindAllGrants()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Contact(ctx context.Context, id string) (*model.ContactDetail, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "_id", Value: id}, {Key: "created_by", Value: user_id}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}}
	}
	res, err := methods.FindContactDetail(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Contacts(ctx context.Context) ([]*model.ContactOverview, error) {
	var filter bson.D
	user_id, err := auth.ForUserID(ctx)
	if err != nil {
		return nil, err
	}
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		filter = bson.D{{Key: "created_by", Value: user_id}}
	} else {
		filter = bson.D{{}}
	}
	res, err := methods.FindContacts(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) UserContacts(ctx context.Context, userID string) ([]*model.ContactOverview, error) {
	admin_err := auth.ForAdmin(ctx)
	if admin_err != nil {
		return nil, admin_err
	}
	res, err := methods.FindUserContacts(userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
