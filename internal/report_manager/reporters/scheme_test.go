package reporters_test

import (
	"net/url"
	"testing"

	"github.com/stjudewashere/seonaut/internal/models"
	"github.com/stjudewashere/seonaut/internal/report_manager/reporter_errors"
	"github.com/stjudewashere/seonaut/internal/report_manager/reporters"
)

// Test the HTTPScheme reporter with a PageReport uses the https scheme.
// The reporter should not report the issue.
func TestHTTPSchemeNoIssues(t *testing.T) {
	// Parse an URL with the https scheme.
	pageURL := "https://example.com"
	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		t.Errorf("Parse URL error: %v", err)
	}

	// Create a PageReport.
	pageReport := &models.PageReport{
		Crawled:    true,
		URL:        pageURL,
		ParsedURL:  parsedURL,
		StatusCode: 200,
	}

	// Create a new HTTPSchemeReporter.
	reporter := reporters.NewHTTPSchemeReporter()
	if reporter.ErrorType != reporter_errors.ErrorHTTPScheme {
		t.Errorf("TestNoIssues: error type is not correct")
	}

	// Run the reporter callback with the PageReport.
	reportsIssue := reporter.Callback(pageReport)

	// The reporter should not found any issue.
	if reportsIssue == true {
		t.Errorf("TestHTTPSchemeNoIssues: reportsIssue should be false")
	}
}

// Test the HTTPScheme reporter with a PageReport that uses the http scheme.
// The reporter should report the issue.
func TestHTTPSchemeIssues(t *testing.T) {
	// Parse an URL with the http scheme
	pageURL := "http://example.com"
	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		t.Errorf("Parse URL error: %v", err)
	}

	// Create a PageReport.
	pageReport := &models.PageReport{
		Crawled:    true,
		URL:        pageURL,
		ParsedURL:  parsedURL,
		StatusCode: 200,
	}

	// Create a new HTTPSchemeReporter.
	reporter := reporters.NewHTTPSchemeReporter()
	if reporter.ErrorType != reporter_errors.ErrorHTTPScheme {
		t.Errorf("TestNoIssues: error type is not correct")
	}

	// Run the reporter callback with the PageReport.
	reportsIssue := reporter.Callback(pageReport)

	// The reporter should found an issue.
	if reportsIssue == false {
		t.Errorf("TestHTTPSchemeIssues: reportsIssue should be true")
	}
}
