package router

import (
	"net/http"
	"regexp"
	"rest/controller"
)

func NewRouter(taskController *controller.TaskController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		re := regexp.MustCompile(`^/tasks/(\d+)$`)
		re1 := regexp.MustCompile(`^/tasks/(\d+)/complete$`)

		if r.RequestURI == "/tasks" && r.Method == http.MethodGet {
			taskController.FindAll(w, r)
		} else if r.RequestURI == "/tasks" && r.Method == http.MethodPost {
			taskController.Create(w, r)
		} else if re.MatchString(r.RequestURI) && r.Method == http.MethodDelete {
			taskController.Delete(w, r)
		} else if re.MatchString(r.RequestURI) && r.Method == http.MethodPut {
			taskController.Update(w, r)
		} else if re1.MatchString(r.RequestURI) && r.Method == http.MethodPatch {
			taskController.CompletedTask(w, r)
		} else {
			w.WriteHeader(404)
		}
	}
}
