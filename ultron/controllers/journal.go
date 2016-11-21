package controllers

import (
	"net/http"
	"strconv"

	"./../models"
	"github.com/gernest/utron"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

/*
	As far as I can see, only one controller I need
	That is Journal
*/
type Journal struct {
	*utron.BaseController
	Routes []string
}

/*
	Home renders the journal which consist of three
	views so far
	- Daily
	- Monthly
	- Future
  However, this is sort of a SPA
*/
func (t *Journal) Home() {
	tasks := []*models.Task{}
	t.Ctx.DB.Order("created_at desc").Find(&tasks)
	t.Ctx.Data["List"] = tasks
	t.Ctx.Template = "journal/index"
	t.HTML(http.StatusOK)
}

/*
	Create a type of task
*/
func (t *Journal) Create() {
	task := &models.Task{}
	request := t.Ctx.Request()
	_ = request.ParseForm()

	if err := decoder.Decode(task, request.PostForm); err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}

	t.Ctx.DB.Create(task)
	t.Ctx.Redirect("/", http.StatusFound)
}

/*
	Delete a type of task
*/
func (t *Journal) Delete() {
	taskID := t.Ctx.Params["id"]
	ID, err := strconv.Atoi(taskID)
	if err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}
	t.Ctx.DB.Delete(&models.Task{ID: ID})
	t.Ctx.Redirect("/", http.StatusFound)
}

/*
	NewJournal returns a new Journal controller
*/
func NewJournal() *Journal {
	return &Journal{
		Routes: []string{
			"get;/api;Home",
			"post;/api/tasks;Create",
			"get;/api/tasks/{id};Delete",
		},
	}
}

/*
	Init a Journal
*/
func init() {
	utron.RegisterController(NewJournal())
}
