package reporters

import (
	"github.com/stjudewashere/seonaut/internal/models"
	"github.com/stjudewashere/seonaut/internal/report_manager"
	"github.com/stjudewashere/seonaut/internal/report_manager/reporter_errors"
)

// Returns a new report_manager.PageIssueReporter with a callback function that
// checks if the status code is in the 30x range.
func NewStatus30xReporter() *report_manager.PageIssueReporter {
	c := func(pageReport *models.PageReport) bool {
		if pageReport.Crawled == false {
			return false
		}

		return pageReport.StatusCode >= 300 && pageReport.StatusCode < 400
	}

	return &report_manager.PageIssueReporter{
		ErrorType: reporter_errors.Error30x,
		Callback:  c,
	}
}

// Returns a new report_manager.PageIssueReporter with a callback function that
// checks if the status code is in the 40x range.
func NewStatus40xReporter() *report_manager.PageIssueReporter {
	c := func(pageReport *models.PageReport) bool {
		if pageReport.Crawled == false {
			return false
		}

		return pageReport.StatusCode >= 400 && pageReport.StatusCode < 500
	}

	return &report_manager.PageIssueReporter{
		ErrorType: reporter_errors.Error40x,
		Callback:  c,
	}
}

// Returns a new report_manager.PageIssueReporter with a callback function that
// checks if the status code is greater or equal than 500.
func NewStatus50xReporter() *report_manager.PageIssueReporter {
	c := func(pageReport *models.PageReport) bool {
		if pageReport.Crawled == false {
			return false
		}

		return pageReport.StatusCode >= 500
	}

	return &report_manager.PageIssueReporter{
		ErrorType: reporter_errors.Error50x,
		Callback:  c,
	}
}
