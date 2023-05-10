package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	generated1 "thomps9012/prevention_productivity/graph/generated"
	"thomps9012/prevention_productivity/graph/model"
)

// isAdmin := auth.ForAdmin(ctx)
// userID := auth.ForUserID(ctx)
func (r *mutationResolver) CreateUser(ctx context.Context, newUser model.NewUser) (*model.LoginRes, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) Login(ctx context.Context, login model.LoginInput) (*model.LoginRes, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, updateUser model.UpdateUser, id string) (*model.UserUpdateRes, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.UserUpdateRes, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) CreateGrant(ctx context.Context, newGrant model.NewGrant) (*model.GrantDetail, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) UpdateGrant(ctx context.Context, id string, updateGrant model.UpdateGrant) (*model.Grant, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) DeleteGrant(ctx context.Context, id string) (bool, error) {
	return false, errors.New("method unimplemented")
}

func (r *mutationResolver) CreateContact(ctx context.Context, newContact model.NewContact) (*model.ContactDetail, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) UpdateContact(ctx context.Context, id string, updateContact model.UpdateContact) (*model.Contact, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) DeleteContact(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateNote(ctx context.Context, newNote model.NewNote) (*model.NoteDetail, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) UpdateNote(ctx context.Context, id string, updateNote model.UpdateNote) (*model.Note, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) DeleteNote(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateLog(ctx context.Context, newLog model.NewLog) (*model.LogRes, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) UpdateLog(ctx context.Context, id string, updateLog model.UpdateLog) (*model.Log, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) DeleteLog(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ApproveLog(ctx context.Context, id string) (bool, error) {
	return false, errors.New("method unimplemented")
}

func (r *mutationResolver) RejectLog(ctx context.Context, id string) (bool, error) {
	return false, errors.New("method unimplemented")
}

func (r *mutationResolver) CreateEvent(ctx context.Context, newEvent model.NewEvent) (*model.EventRes, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, id string, updateEvent model.UpdateEvent) (*model.Event, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) DeleteEvent(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ApproveEvent(ctx context.Context, id string) (bool, error) {
	return false, errors.New("method unimplemented")
}

func (r *mutationResolver) RejectEvent(ctx context.Context, id string) (bool, error) {
	return false, errors.New("method unimplemented")
}

func (r *mutationResolver) CreateEventSummary(ctx context.Context, newEventSummary model.NewEventSummary) (*model.EventSummaryRes, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) UpdateEventSummary(ctx context.Context, id string, updateEventSummary model.UpdateEventSummary) (*model.EventSummary, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) DeleteEventSummary(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ApproveEventSummary(ctx context.Context, id string) (bool, error) {
	return false, errors.New("method unimplemented")
}

func (r *mutationResolver) RejectEventSummary(ctx context.Context, id string) (bool, error) {
	return false, errors.New("method unimplemented")
}

func (r *mutationResolver) CreateSchoolReportPlan(ctx context.Context, newSchoolReportPlan model.NewSchoolReportPlan) (*model.SchoolReportPlanRes, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) UpdateSchoolReportPlan(ctx context.Context, id string, updateSchoolReportPlan model.UpdateSchoolReportPlan) (*model.SchoolReportPlan, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) DeleteSchoolReportPlan(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ApproveSchoolReportPlan(ctx context.Context, id string) (bool, error) {
	return false, errors.New("method unimplemented")
}

func (r *mutationResolver) RejectSchoolReportPlan(ctx context.Context, id string) (bool, error) {
	return false, errors.New("method unimplemented")
}

func (r *mutationResolver) CreateSchoolReportDebrief(ctx context.Context, newSchoolReportDebrief model.NewSchoolReportDebrief) (*model.SchoolReportDebriefRes, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) UpdateSchoolReportDebrief(ctx context.Context, id string, updateSchoolReportDebrief model.UpdateSchoolReportDebrief) (*model.SchoolReportDebrief, error) {
	return nil, errors.New("method unimplemented")
}

func (r *mutationResolver) DeleteSchoolReportDebrief(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ApproveSchoolReportDebrief(ctx context.Context, id string) (bool, error) {
	return false, errors.New("method unimplemented")
}

func (r *mutationResolver) RejectSchoolReportDebrief(ctx context.Context, id string) (bool, error) {
	return false, errors.New("method unimplemented")
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return nil, errors.New("method unimplemented")
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return nil, errors.New("method unimplemented")
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return nil, errors.New("method unimplemented")
}

func (r *queryResolver) ItemNotes(ctx context.Context, itemID string) ([]*model.NoteDetail, error) {
	return nil, errors.New("method unimplemented")
}

func (r *queryResolver) Note(ctx context.Context, id string) (*model.NoteDetail, error) {
	return nil, errors.New("method unimplemented")
}

func (r *queryResolver) UserNotes(ctx context.Context, userID string) ([]*model.Note, error) {
	return nil, errors.New("method unimplemented")
}

func (r *queryResolver) Log(ctx context.Context, id string) (*model.LogWithNotes, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) AllLogs(ctx context.Context) ([]*model.LogOverview, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) UserLogs(ctx context.Context, userID string) ([]*model.LogOverview, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) Event(ctx context.Context, id string) (*model.EventWithNotes, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) Events(ctx context.Context) ([]*model.EventOverview, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) UserEvents(ctx context.Context, userID string) ([]*model.EventOverview, error) {
	return nil, errors.New("method unimplemented")
}

func (r *queryResolver) EventSummary(ctx context.Context, id string) (*model.EventSummaryWithNotes, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) EventSummaries(ctx context.Context) ([]*model.EventSummaryOverview, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) UserEventSummaries(ctx context.Context, userID string) ([]*model.EventSummaryOverview, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) SchoolReportPlan(ctx context.Context, id string) (*model.SchoolReportPlanWithNotes, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) SchoolReportPlans(ctx context.Context) ([]*model.SchoolReportPlanOverview, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) UserSchoolReportPlans(ctx context.Context, userID string) ([]*model.SchoolReportPlanOverview, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) SchoolReportDebrief(ctx context.Context, id string) (*model.SchoolReportDebriefWithNotes, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) SchoolReportDebriefs(ctx context.Context) ([]*model.SchoolReportDebriefOverview, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) UserSchoolReportDebriefs(ctx context.Context, userID string) ([]*model.SchoolReportDebriefOverview, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) Grant(ctx context.Context, id string) (*model.GrantDetail, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) Grants(ctx context.Context) ([]*model.GrantOverview, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) Contact(ctx context.Context, id string) (*model.ContactDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Contacts(ctx context.Context) ([]*model.ContactOverview, error) {
	return nil, errors.New("method unimplemented")

}

func (r *queryResolver) UserContacts(ctx context.Context) ([]*model.ContactOverview, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
