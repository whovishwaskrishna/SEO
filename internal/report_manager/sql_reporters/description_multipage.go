package sql_reporters

import (
	"github.com/stjudewashere/seonaut/internal/models"
	"github.com/stjudewashere/seonaut/internal/report_manager"
	"github.com/stjudewashere/seonaut/internal/report_manager/reporter_errors"
)

// Creates a MultipageIssueReporter object that contains the SQL query to check for pages
// that have the same description, taking into account the status code and language.
func (sr *SqlReporter) DuplicatedDescriptionReporter(c *models.Crawl) *report_manager.MultipageIssueReporter {
	query := `
		SELECT
			y.id
		FROM pagereports y
		INNER JOIN (
			SELECT
				description,
				lang,
				count(*) AS c
			FROM pagereports
			WHERE crawl_id = ? AND media_type = "text/html" AND status_code >= 200
			AND status_code < 300 AND (canonical = "" OR canonical = url) AND crawled = 1
			GROUP BY description, lang
			HAVING c > 1
		) d 
		ON d.description = y.description AND d.lang = y.lang
		WHERE y.media_type = "text/html" AND length(y.description) > 0 AND y.crawl_id = ?
		AND status_code >= 200 AND status_code < 300 AND (canonical = "" OR canonical = url AND crawled = 1)`

	return &report_manager.MultipageIssueReporter{
		Pstream:   sr.pageReportsQuery(query, c.Id, c.Id),
		ErrorType: reporter_errors.ErrorDuplicatedDescription,
	}
}
