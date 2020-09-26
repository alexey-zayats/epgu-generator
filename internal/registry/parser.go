package registry

import (
	"epgu-generator/internal/model"
	"epgu-generator/internal/util"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strings"
)

// Parser ...
type Parser struct {
	data map[string][]*model.Registry
}

// NewParser ...
func NewParser() *Parser {
	return &Parser{
		data: make(map[string][]*model.Registry),
	}
}

// Parse ...
func (p *Parser) Parse(filePath string) (map[string][]*model.Registry, error) {

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return p.data, errors.Wrapf(err, "unable open xlsx file %s", filePath)
	}

	for _, sheet := range f.GetSheetMap() {
		logrus.Debug(sheet)

		rows, err := f.Rows(sheet)
		if err != nil {
			return p.data, errors.Wrapf(err, "unable get rows by sheet %s", sheet)
		}

		i := 2
		for rows.Next() {

			i++

			axis := map[string]string{
				"departmentName":    fmt.Sprintf("A%d", i),
				"departmentCode":    fmt.Sprintf("B%d", i),
				"serviceName":       fmt.Sprintf("C%d", i),
				"serviceTargetName": fmt.Sprintf("D%d", i),
				"serviceTargetID":   fmt.Sprintf("E%d", i),
				"serviceFormCode":   fmt.Sprintf("G%d", i),
				"applicantType":     fmt.Sprintf("H%d", i),
				"useSignature":      fmt.Sprintf("I%d", i),
				"unlinkService":     fmt.Sprintf("J%d", i),
				"change":            fmt.Sprintf("K%d", i),
			}

			departmentName := f.GetCellValue(sheet, axis["departmentName"])
			departmentCode := f.GetCellValue(sheet, axis["departmentCode"])
			serviceName := f.GetCellValue(sheet, axis["serviceName"])
			serviceTargetName := f.GetCellValue(sheet, axis["serviceTargetName"])
			serviceTargetID := f.GetCellValue(sheet, axis["serviceTargetID"])
			serviceFormCode := f.GetCellValue(sheet, axis["serviceFormCode"])
			applicantType := f.GetCellValue(sheet, axis["applicantType"])
			useSignature := f.GetCellValue(sheet, axis["useSignature"])

			unlinkService := f.GetCellValue(sheet, axis["unlinkService"])
			change := f.GetCellValue(sheet, axis["change"])

			if serviceTargetID == "" {
				break
			}

			r := &model.Registry{
				DepartmentName:    departmentName,
				DepartmentCode:    departmentCode,
				ServiceName:       serviceName,
				ServiceTargetName: serviceTargetName,
				ServiceTargetID:   serviceTargetID,
				ServiceFormCode:   serviceFormCode,
				ApplicantType:     util.ParseApplicant(applicantType),
				UseSignature:      util.ParseUseSignature(useSignature),
				UnlinkService:     strings.Split(unlinkService, ","),
				Change:            change,
			}

			v, ok := p.data[r.ServiceFormCode]
			if ok == false {
				v = make([]*model.Registry, 0)
			}

			v = append(v, r)
			p.data[r.ServiceFormCode] = v
		}
	}

	return p.data, nil
}
