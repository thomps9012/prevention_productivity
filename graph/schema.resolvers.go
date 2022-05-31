package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"prevention_productivity/graph/generated"
	"prevention_productivity/graph/model"
	"prevention_productivity/db"
	"time"
)

func (r *mutationResolver) AddUser(ctx context.Context, input *model.UserInput) (*model.User, error) {
	user :=  db.C("users").Insert(input)
	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *model.UserInput) (*model.User, error) {
	user := ctx.Value("user")
	if user.admin == true {
		db.C("users").UpdateId(id, input)
		} else if user._id == id {
			db.C("users").UpdateId(id, input)
			} else {
				panic(fmt.Errorf("not authorized"))
			}
			return &user, nil
		}

		
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	deletedUser := db.C("users").RemoveId(id)
	return deletedUser, nil
}

// func (r *mutationResolver) AddGrant(ctx context.Context, input *model.GrantInput) (*model.Grant, error) {
// 	var user = ctx.Value("user")
// 	if user.admin == true {
// 		var grant = db.C("grants").Insert(input)
// 	} else {
// 		panic(fmt.Errorf("not authorized"))
// 	}
// 	return grant, nil
// }

// func (r *mutationResolver) UpdateGrant(ctx context.Context, id string, input *model.GrantInput) (*model.Grant, error) {
// 	var user = ctx.Value("user")
// 	if user.admin == true {
// 		var grant = db.C("grants").UpdateId(id, input)
// 	} else {
// 		panic(fmt.Errorf("not authorized"))
// 	}
// 	return grant, nil
// }

// func (r *mutationResolver) DeleteGrant(ctx context.Context, id string) (*model.Grant, error) {
// 	var user = ctx.Value("user")
// 	if user.admin == true {
// 		db.C("grants").RemoveId(id)
// 	} else {
// 		panic(fmt.Errorf("not authorized"))
// 	}
// 	return nil, nil
// }

// func (r *mutationResolver) AddProductivityLog(ctx context.Context, input *model.ProductivityLogInput) (*model.ProductivityLog, error) {
// 	var user = ctx.Value("user")
// 	log := &model.ProductivityLog{
// 		date:         input.Date,
// 		grant:        input.Grant,
// 		focus_area:   input.FocusArea,
// 		actions:      input.Actions,
// 		successes:    input.Successes,
// 		improvements: input.Improvements,
// 		next_steps:   input.NextSteps,
// 		status:       "PENDING",
// 		author:       user,
// 	}
// 	db.C("productivity_logs").Insert(log)
// 	return &log, nil
// }

// func (r *mutationResolver) UpdateProductivityLog(ctx context.Context, id string, input *model.ProductivityLogInput) (*model.ProductivityLog, error) {
// 	var user = ctx.Value("user")
// 	var log = db.C("productivity_logs").FindId(id)
// 	if user.admin == true {
// 		db.C("productivity_logs").UpdateId(id, input)
// 	} else if user.id == log.author.id {
// 		db.C("productivity_logs").UpdateId(id, input)
// 	} else {
// 		panic(fmt.Errorf("not authorized"))
// 	}
// 	return log, nil
// }

// func (r *mutationResolver) DeleteProductivityLog(ctx context.Context, id string) (*model.ProductivityLog, error) {
// 	var user = ctx.Value("user")
// 	if user.admin == true {
// 		db.C("productivity_logs").RemoveId(id)
// 	} else {
// 		panic(fmt.Errorf("not authorized"))
// 	}
// 	return nil, nil
// }

// func (r *mutationResolver) AddEvent(ctx context.Context, input *model.EventInput) (*model.Event, error) {
// 	event := Event{}
// 	event.status = "PENDING"
// 	event.event_lead = ctx.Value("user")
// 	db.C("events").Insert(event)
// 	return &event, nil
// }

// func (r *mutationResolver) UpdateEvent(ctx context.Context, id string, input *model.EventInput) (*model.Event, error) {
// 	var user = ctx.Value("user")
// 	if user.admin == true {
// 		db.C("events").UpdateId(id, input)
// 	} else if user.id == id {
// 		db.C("events").UpdateId(id, input)
// 	} else {
// 		panic(fmt.Errorf("not authorized"))
// 	}
// 	return &event, nil
// }

// func (r *mutationResolver) DeleteEvent(ctx context.Context, id string) (*model.Event, error) {
// 	var user = ctx.Value("user")
// 	if user.admin == true {
// 		db.C("events").RemoveId(id)
// 	} else {
// 		panic(fmt.Errorf("not authorized"))
// 	}
// 	return nil, nil
// }

// func (r *mutationResolver) AddEventSummary(ctx context.Context, input *model.EventSummaryInput) (*model.EventSummary, error) {
// 	var user = ctx.Value("user")
// 	var today = time.Now()
// 	summary := &model.EventSummary{
// 		date:       input.date,
// 		event:      input.event,
// 		attendees:  input.attendees,
// 		challenges: input.challenges,
// 		outcomes:   input.outcomes,
// 		next_steps: input.next_steps,
// 		status:     "PENDING",
// 		created_at: today,
// 		author:     user,
// 	}
// 	// summary.status = "PENDING"
// 	// summary.created_at = time.Now()
// 	// summary.author = ctx.Value("user")
// 	db.C("event_summaries").Insert(summary)
// 	return &summary, nil
// }

// func (r *mutationResolver) UpdateEventSummary(ctx context.Context, id string, input *model.EventSummaryInput) (*model.EventSummary, error) {
// 	var user = ctx.Value("user")
// 	var summary = db.C("event_summaries").FindId(id)
// 	if user.admin == true {
// 		db.C("event_summaries").UpdateId(id, input)
// 	} else if user.id == summary.author.id {
// 		db.C("event_summaries").UpdateId(id, input)
// 	} else {
// 		panic(fmt.Errorf("not authorized"))
// 	}
// 	return &summary, nil
// }

// func (r *mutationResolver) DeleteEventSummary(ctx context.Context, id string) (*model.EventSummary, error) {
// 	var user = ctx.Value("user")
// 	if user.admin == true {
// 		db.C("event_summaries").RemoveId(id)
// 	} else {
// 		panic(fmt.Errorf("not authorized"))
// 	}
// 	return &summary, nil
// }

// func (r *mutationResolver) AddSchoolReport(ctx context.Context, input *model.SchoolReportInput) (*model.SchoolReport, error) {
// 	var user = ctx.Value("user")
// 	report := SchoolReport{}
// 	report.status = "PENDING"
// 	report.author = user
// 	db.C("school_reports").Insert(report)
// 	return &report, nil
// }

// func (r *mutationResolver) UpdateSchoolReport(ctx context.Context, id string, input *model.SchoolReportInput) (*model.SchoolReport, error) {
// 	var user = ctx.Value("user")
// 	var report = db.C("school_reports").FindId(id)
// 	if user.admin == true {
// 		db.C("school_reports").UpdateId(id, input)
// 	} else if user.id == report.author.id {
// 		db.C("school_reports").UpdateId(id, input)
// 	} else {
// 		panic(fmt.Errorf("not authorized"))
// 	}
// 	return &report, nil
// }

// func (r *mutationResolver) DeleteSchoolReport(ctx context.Context, id string) (*model.SchoolReport, error) {
// 	var user = ctx.Value("user")
// 	if user.admin == true {
// 		db.C("school_reports").RemoveId(id)
// 	} else {
// 		panic(fmt.Errorf("not authorized"))
// 	}
// 	return &report, nil
// }

// func (r *mutationResolver) AddContact(ctx context.Context, input *model.ContactInput) (*model.Contact, error) {
// 	var contact = db.C("contacts").Insert(input)
// 	return &contact, nil
// }

// func (r *mutationResolver) UpdateContact(ctx context.Context, id string, input *model.ContactInput) (*model.Contact, error) {
// 	var user = ctx.Value("user")
// 	if user.admin == true {
// 		db.C("contacts").UpdateId(id, input)
// 	} else {
// 		panic(fmt.Errorf("not authorized"))
// 	}
// 	return &contact, nil
// }

// func (r *mutationResolver) DeleteContact(ctx context.Context, id string) (*model.Contact, error) {
// 	var user = ctx.Value("user")
// 	if user.admin == true {
// 		var contact = db.C("contacts").RemoveId(id)
// 	} else {
// 		panic(fmt.Errorf("not authorized"))
// 	}
// 	return contact, nil
// }

// func (r *mutationResolver) AddNote(ctx context.Context, input *model.NoteInput) (*model.Note, error) {
// 	user := ctx.Value("user")
// 	var note = &model.Note{
// 		text:    input.text,
// 		section: input.section,
// 		item_id: input.item_id,
// 	}
// 	note.Date = time.Now()
// 	note.Author = user
// 	db.C("notes").Insert(note)
// 	return &note, nil
// }

// func (r *mutationResolver) EditNote(ctx context.Context, id string, input *model.NoteInput) (*model.Note, error) {
// 	var user = ctx.Value("user")
// 	var note = db.C("notes").FindId(id)
// 	if user.id == note.author.id {
// 		db.C("notes").UpdateId(id, input)
// 	} else {
// 		panic(fmt.Errorf("not authorized"))
// 	}
// 	return &note, nil
// }

// func (r *noteResolver) ID(ctx context.Context, obj *model.Note) (string, error) {
// 	panic(fmt.Errorf("not implemented"))
// }

// func (r *productivityLogResolver) ApprovalStatus(ctx context.Context, obj *model.ProductivityLog) (model.ApprovalStatus, error) {
// 	panic(fmt.Errorf("not implemented"))
// }

// func (r *productivityLogResolver) Notes(ctx context.Context, obj *model.ProductivityLog) ([]*model.Note, error) {
// 	panic(fmt.Errorf("not implemented"))
// }

// func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
// 	users := []*model.User{}
// 	db.C("users").Find(bson.M{}).All(&users)
// 	return users, nil
// }

// func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
// 	user := db.C("users").FindId(id)
// 	return &user, nil
// }

// func (r *queryResolver) Grants(ctx context.Context) ([]*model.Grant, error) {
// 	grants := []*model.Grant{}
// 	db.C("grants").Find(bson.M{}).All(&grants)
// 	return grants, nil
// }

// func (r *queryResolver) Grant(ctx context.Context, id string) (*model.Grant, error) {
// 	grant := db.C("grants").FindId(id)
// 	return &grant, nil
// }

// func (r *queryResolver) ProductivityLogs(ctx context.Context) ([]*model.ProductivityLog, error) {
// 	productivityLogs := []*model.ProductivityLog{}
// 	db.C("productivity_logs").Find(bson.M{}).All(&productivityLogs)
// 	return productivityLogs, nil
// }

// func (r *queryResolver) ProductivityLog(ctx context.Context, id string) (*model.ProductivityLog, error) {
// 	productivityLog := db.C("productivity_logs").FindId(id)
// 	return &productivityLog, nil
// }

// func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
// 	events := []*model.Event{}
// 	db.C("events").Find(bson.M{}).All(&events)
// 	return events, nil
// }

// func (r *queryResolver) Event(ctx context.Context, id string) (*model.Event, error) {
// 	event := db.C("events").FindId(id)
// 	return &event, nil
// }

// func (r *queryResolver) EventSummaries(ctx context.Context) ([]*model.EventSummary, error) {
// 	eventSummaries := []*model.EventSummary{}
// 	db.C("event_summaries").Find(bson.M{}).All(&eventSummaries)
// 	return eventSummaries, nil
// }

// func (r *queryResolver) EventSummary(ctx context.Context, id string) (*model.EventSummary, error) {
// 	eventSummary := db.C("event_summaries").FindId(id)
// 	return &eventSummary, nil
// }

// func (r *queryResolver) SchoolReports(ctx context.Context) ([]*model.SchoolReport, error) {
// 	schoolReports := []*model.SchoolReport{}
// 	db.C("school_reports").Find(bson.M{}).All(&schoolReports)
// 	return schoolReports, nil
// }

// func (r *queryResolver) SchoolReport(ctx context.Context, id string) (*model.SchoolReport, error) {
// 	schoolReport := db.C("school_reports").FindId(id)
// 	return &schoolReport, nil
// }

// func (r *queryResolver) Contacts(ctx context.Context) ([]*model.Contact, error) {
// 	contacts := []*model.Contact{}
// 	db.C("contacts").Find(bson.M{}).All(&contacts)
// 	return contacts, nil
// }

// func (r *queryResolver) Contact(ctx context.Context, id string) (*model.Contact, error) {
// 	contact := db.C("contacts").FindId(id)
// 	return &contact, nil
// }

// func (r *schoolReportResolver) ApprovalStatus(ctx context.Context, obj *model.SchoolReport) (model.ApprovalStatus, error) {
// 	panic(fmt.Errorf("not implemented"))
// }

// func (r *schoolReportResolver) Notes(ctx context.Context, obj *model.SchoolReport) ([]*model.Note, error) {
// 	panic(fmt.Errorf("not implemented"))
// }

// func (r *schoolReportResolver) CreatedAt(ctx context.Context, obj *model.SchoolReport) (string, error) {
// 	panic(fmt.Errorf("not implemented"))
// }

// Contact returns generated.ContactResolver implementation.
func (r *Resolver) Contact() generated.ContactResolver { return &contactResolver{r} }

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

// EventSummary returns generated.EventSummaryResolver implementation.
func (r *Resolver) EventSummary() generated.EventSummaryResolver { return &eventSummaryResolver{r} }

// Grant returns generated.GrantResolver implementation.
func (r *Resolver) Grant() generated.GrantResolver { return &grantResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Note returns generated.NoteResolver implementation.
func (r *Resolver) Note() generated.NoteResolver { return &noteResolver{r} }

// ProductivityLog returns generated.ProductivityLogResolver implementation.
func (r *Resolver) ProductivityLog() generated.ProductivityLogResolver {
	return &productivityLogResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// SchoolReport returns generated.SchoolReportResolver implementation.
func (r *Resolver) SchoolReport() generated.SchoolReportResolver { return &schoolReportResolver{r} }

type contactResolver struct{ *Resolver }
type eventResolver struct{ *Resolver }
type eventSummaryResolver struct{ *Resolver }
type grantResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type noteResolver struct{ *Resolver }
type productivityLogResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type schoolReportResolver struct{ *Resolver }
