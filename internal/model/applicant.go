package model

// ApplicantType ...
type ApplicantType int

const (
	// ApplicantNone ...
	ApplicantNone ApplicantType = iota

	// ApplicantNP физическое лицо (Natural Person)
	ApplicantNP

	// ApplicantIE индивидуальный предпрениматель (Individual Entrepreneur)
	ApplicantIE

	// ApplicantLE юридическое лицо (Legal Entity)
	ApplicantLE
)
