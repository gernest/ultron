package controllers

import (
	"net/http"
	"strconv"

	"github.com/gernest/utron/controller"
	"github.com/gorilla/schema"
	"github.com/thelastinuit/ultron/models"
)

var decoder = schema.NewDecoder()

type Journal struct {
	controller.BaseController
	Routes []string
}

func (t *Journal) Home() {
	t.Ctx.Template = "journal/index"
	t.HTML(http.StatusOK)
}

func (t *Journal) Index() {
	tasks := []*models.Task{}
	t.Ctx.DB.Order("created_at desc").Find(&tasks)
	t.Ctx.Data["data"] = tasks
	t.JSON(200)
}

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

func NewJournal() controller.Controller {
	return &Journal{
		Routes: []string{
			"get;/;Home",
			"get;/api/tasks;Index",
			"post;/api/tasks;Create",
			"delete;/api/tasks/{id};Delete",
		},
	}
}
