package sql_reporters

import (
	"database/sql"
	"log"

	"github.com/stjudewashere/seonaut/internal/report_manager"
)

type SqlReporter struct {
	db *sql.DB
}

// NewSqlReporter creates a new SqlReporter with the given SQL database connection.
func NewSqlReporter(db *sql.DB) *SqlReporter {
	return &SqlReporter{
		db: db,
	}
}

// GetAllReporters returns a slice of all the implemented reporters in the SqlReporter.
func (sr *SqlReporter) GetAllReporters() []report_manager.MultipageCallback {
	return []report_manager.MultipageCallback{
		// Add status code issue reporters
		sr.RedirectChainsReporter,
		sr.RedirectLoopsReporter,

		// Add title issue reporters
		sr.DuplicatedTitleReporter,

		// Add description issue reporters
		sr.DuplicatedDescriptionReporter,

		// Add link issue reporters
		sr.OrphanPagesReporter,
		sr.NoFollowIndexableReporter,
		sr.FollowNoFollowReporter,

		// Add hreflang reporters
		sr.MissingHrelangReturnLinks,
		sr.HreflangsToNonCanonical,
		sr.HreflangNoindexable,

		// Add canonical issue reporters
		sr.CanonicalizedToNonCanonical,
		sr.CanonicalizedToNonIndexable,
	}
}

// pageReportsQuery executes a SQL query and returns a channel of int64 which is used to send
// the PageReport ids through.
func (sr *SqlReporter) pageReportsQuery(query string, args ...interface{}) <-chan int64 {
	prStream := make(chan int64)

	go func() {
		defer close(prStream)

		rows, err := sr.db.Query(query, args...)
		if err != nil {
			log.Println(err)
		}

		for rows.Next() {
			var pid int64
			err := rows.Scan(&pid)
			if err != nil {
				log.Println(err)
				continue
			}

			prStream <- pid
		}
	}()

	return prStream
}
