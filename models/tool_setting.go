package models

type ToolSetting struct {
	ResourceTypeCode string `json:"resource_type_code"` // the resource type code of the resource handler to use to display originality reports.Example: originality_reports
	ResourceUrl      string `json:"resource_url"`       // a URL that may be used to override the launch URL inferred by the specified resource_type_code. If used a 'resource_type_code' must also be specified..Example: http://www.test.com/originality_report
}

func (t *ToolSetting) HasError() error {
	return nil
}
