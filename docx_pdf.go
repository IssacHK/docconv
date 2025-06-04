package docconv

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// ConvertDocxToPDF converts a DOCX file at path to a PDF inside outDir using libreoffice.
// It returns the path to the generated PDF file.
func ConvertDocxToPDF(path string, outDir string) (string, error) {
	cmd := exec.Command("libreoffice", "--headless", "--convert-to", "pdf", path, "--outdir", outDir)
	if out, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("libreoffice convert: %v, output: %s", err, out)
	}
	base := filepath.Base(path)
	base = strings.TrimSuffix(base, filepath.Ext(base))
	pdfPath := filepath.Join(outDir, base+".pdf")
	return pdfPath, nil
}
