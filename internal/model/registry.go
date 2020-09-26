package model

// Registry ...
type Registry struct {
	DepartmentName    string
	DepartmentCode    string
	ServiceName       string
	ServiceTargetName string
	ServiceTargetID   string
	ServiceFormCode   string
	ApplicantType     []ApplicantType
	UseSignature      bool
	UnlinkService     []string
	Change            string
}
