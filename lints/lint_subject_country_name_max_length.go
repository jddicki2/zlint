// lint_subject_country_name_max_length.go
/************************************************
RFC 5280: A.1
	* In this Appendix, there is a list of upperbounds
	for fields in a x509 Certificate. *
	ub-country-name-alpha-length INTEGER ::= 2
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subjectCountryNameMaxLength struct {
	// Internal data here
}

func (l *subjectCountryNameMaxLength) Initialize() error {
	return nil
}

func (l *subjectCountryNameMaxLength) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return true
}

func (l *subjectCountryNameMaxLength) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, j := range c.Subject.Country {
		if len(j) > 2 {
			return ResultStruct{Result: Error}, nil
		}
	}

	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_country_name_max_length",
		Description:   "The 'Country' field must be less than or equal to 2 characters.",
		Providence:    "RFC 5280: A.1",
		EffectiveDate: util.RFC2459Date,
		Test:          &subjectCountryNameMaxLength{}})
}
