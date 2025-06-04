package docx_test

import (
	"os"
	"path/filepath"
	"testing"

	"code.sajari.com/docconv/v2"
)

func TestConvertDocxToPDF(t *testing.T) {
	tmpDir := t.TempDir()
	pdfPath, err := docconv.ConvertDocxToPDF(filepath.Join("./testdata", "sample.docx"), tmpDir)
	if err != nil {
		t.Fatalf("got error = %v, want nil", err)
	}
	if _, err := os.Stat(pdfPath); err != nil {
		t.Fatalf("expected output %s to exist: %v", pdfPath, err)
	}
}
