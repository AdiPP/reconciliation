package processor

import (
	"testing"

	"github.com/AdiPP/recon-general/mapper"
)

func TestSuccessProcessCSV(t *testing.T) {
	source_path := "../source_test/valid_source_test.csv"
	proxy_path := "../source_test/valid_proxy_test.csv"

	err := ProcessCSV(source_path, proxy_path, "../")

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestFailedProcessCSVCausedByInvalidFile(t *testing.T) {
	source_path := "../source_test/invalid_source_test.csv"
	proxy_path := "../source_test/invalid_proxy_test.csv"

	err := ProcessCSV(source_path, proxy_path, "../")
	expectation := mapper.FILE_NOT_SUPPORTED

	if err == nil {
		t.Errorf("Expected erorr %v but got %v", expectation, err.Error())
	}
}
