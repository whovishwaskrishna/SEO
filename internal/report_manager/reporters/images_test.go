package reporters_test

import (
	"testing"

	"github.com/stjudewashere/seonaut/internal/models"
	"github.com/stjudewashere/seonaut/internal/report_manager/reporter_errors"
	"github.com/stjudewashere/seonaut/internal/report_manager/reporters"
)

// Test the AltText reporter with a pageReport that does not
// have any image without Alt text. The reporter should not report the issue.
func TestAltTextReporterNoIssues(t *testing.T) {
	pageReport := &models.PageReport{
		Crawled:   true,
		MediaType: "text/html",
	}

	// Add an image with alt text
	pageReport.Images = append(pageReport.Images, models.Image{
		Alt: "Image alt text",
	})

	reporter := reporters.NewAltTextReporter()
	if reporter.ErrorType != reporter_errors.ErrorImagesWithNoAlt {
		t.Errorf("TestNoIssues: error type is not correct")
	}

	reportsIssue := reporter.Callback(pageReport)

	if reportsIssue == true {
		t.Errorf("TestAltTextReporterNoIssues: reportsIssue should be false")
	}
}

// Test the LittleContent reporter with a pageReport that does
// have a little content issue. The reporter should report the issue.
func TestAltTextReporterIssues(t *testing.T) {
	pageReport := &models.PageReport{
		Crawled:   true,
		MediaType: "text/html",
	}

	// Add an image without alt text
	pageReport.Images = append(pageReport.Images, models.Image{})

	reporter := reporters.NewAltTextReporter()
	if reporter.ErrorType != reporter_errors.ErrorImagesWithNoAlt {
		t.Errorf("TestNoIssues: error type is not correct")
	}

	reportsIssue := reporter.Callback(pageReport)

	if reportsIssue == false {
		t.Errorf("TestAltTextReporterIssues: reportsIssue should be true")
	}
}
