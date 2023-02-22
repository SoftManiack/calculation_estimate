package service

import (
	dto "calculation-estimate"
	"os"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
)

func RunPdfReport(inquirer dto.Inquirer) ([]dto.Estimate, error) {

	// 1 Grop

	path, _ := os.LookupEnv("PDF")
	pathSources, _ := os.LookupEnv("PATH_SOURCES")

	pdf := gofpdf.New(gofpdf.OrientationPortrait, "pt", "Letter", pathSources)
	pdf.AddFont("Helvetica-1251", "", "helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica-1251", "", 13)
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")

	pdf.CellFormat(0, 20, tr(inquirer.Groups[0].Title),
		"", 1, "", false, 0, "")

	for i := 0; i < len(inquirer.Groups[0].Questions); i++ {

		//h := 50 + 46*float64(i)
		pdf.CellFormat(0, 20, tr(inquirer.Groups[0].Questions[i].Text+" : "+inquirer.Groups[0].Questions[i].Answer),
			"", 1, "", false, 0, "")

	}

	uuid1 := uuid.New()

	err := pdf.OutputFileAndClose(path + uuid1.String() + ".pdf")

	estimate1 := dto.Estimate{
		Uid:  uuid1.String(),
		Name: "Смета1",
	}
	// 2 Grop

	pdf = gofpdf.New(gofpdf.OrientationPortrait, "pt", "Letter", pathSources)
	pdf.AddFont("Helvetica-1251", "", "helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica-1251", "", 13)

	pdf.CellFormat(0, 20, tr(inquirer.Groups[1].Title),
		"", 1, "", false, 0, "")

	for i := 0; i < len(inquirer.Groups[1].Questions); i++ {

		//h := 50 + 46*float64(i)
		pdf.CellFormat(0, 20, tr(inquirer.Groups[1].Questions[i].Text+" : "+inquirer.Groups[1].Questions[i].Answer),
			"", 1, "", false, 0, "")

	}

	uuid2 := uuid.New()

	err = pdf.OutputFileAndClose(path + uuid2.String() + ".pdf")

	estimate2 := dto.Estimate{
		Uid:  uuid2.String(),
		Name: "Смета2",
	}

	var estimates = []dto.Estimate{
		estimate1, estimate2,
	}

	return estimates, err
}
