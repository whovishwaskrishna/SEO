package http

import (
	"net/http"
	"strconv"

	"github.com/stjudewashere/seonaut/internal/models"
	"github.com/stjudewashere/seonaut/internal/projectview"
)

type ExplorerView struct {
	ProjectView   *projectview.ProjectView
	Term          string
	PaginatorView models.PaginatorView
}

// handleExplorer handles the URL explorer request.
// It performas a search of pagereports based on the "term" parameter. In case the "term" parameter
// is empty, it loads all the pagereports.
// It expects a query parameter "pid" containing the project ID, the "p" parameter containing the current
// page in the paginator, and the "term" parameter used to perform the pagereport search.
func (app *App) handleExplorer(w http.ResponseWriter, r *http.Request) {
	// Get user from the request's context
	user, ok := app.userService.GetUserFromContext(r.Context())
	if ok == false {
		http.Redirect(w, r, "/signout", http.StatusSeeOther)
		return
	}

	// Get the project id
	pid, err := strconv.Atoi(r.URL.Query().Get("pid"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Get the page number and set page number to 1 if the parameter is not set
	page, err := strconv.Atoi(r.URL.Query().Get("p"))
	if err != nil {
		page = 1
	}

	// Get the project view
	pv, err := app.projectViewService.GetProjectView(pid, user.Id)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	term := r.URL.Query().Get("term")

	// Get the paginated reports
	paginatorView, err := app.reportService.GetPaginatedReports(pv.Crawl.Id, page, term)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	view := ExplorerView{
		ProjectView:   pv,
		Term:          term,
		PaginatorView: paginatorView,
	}

	v := &PageView{
		Data:      view,
		User:      *user,
		PageTitle: "EXPLORER",
	}

	app.renderer.RenderTemplate(w, "explorer", v)
}
