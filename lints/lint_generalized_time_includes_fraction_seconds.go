// lint_generalized_time_includes_fraction_seconds.go
/********************************************************************
4.1.2.5.2.  GeneralizedTime
The generalized time type, GeneralizedTime, is a standard ASN.1 type
for variable precision representation of time.  Optionally, the
GeneralizedTime field can include a representation of the time
differential between local and Greenwich Mean Time.

For the purposes of this profile, GeneralizedTime values MUST be
expressed in Greenwich Mean Time (Zulu) and MUST include seconds
(i.e., times are YYYYMMDDHHMMSSZ), even where the number of seconds
is zero.  GeneralizedTime values MUST NOT include fractional seconds.
********************************************************************/

package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type generalizedTimeFraction struct {
	date1Gen bool
	date2Gen bool
}

func (l *generalizedTimeFraction) Initialize() error {
	return nil
}

func (l *generalizedTimeFraction) CheckApplies(c *x509.Certificate) bool {
	firstDate, secondDate := util.GetTimes(c)
	beforeTag, afterTag := util.FindTimeType(firstDate, secondDate)
	l.date1Gen = beforeTag == 24
	l.date2Gen = afterTag == 24
	return l.date1Gen || l.date2Gen
}

func (l *generalizedTimeFraction) RunTest(c *x509.Certificate) (ResultStruct, error) {
	r := Pass
	date1, date2 := util.GetTimes(c)
	if l.date1Gen {
		// UTC Tests on notBefore
		checkFraction(&r, date1)
		if r == Error {
			return ResultStruct{Result: r}, nil
		}
	}
	if l.date2Gen {
		checkFraction(&r, date2)
	}
	return ResultStruct{Result: r}, nil
}

func checkFraction(r *ResultEnum, t asn1.RawValue) {
	if t.Bytes[len(t.Bytes)-1] == 'Z' {
		if len(t.Bytes) > 15 {
			*r = Error
		}
	} else if t.Bytes[len(t.Bytes)-5] == '-' || t.Bytes[len(t.Bytes)-1] == '+' {
		if len(t.Bytes) > 19 {
			*r = Error
		}
	} else {
		if len(t.Bytes) > 14 {
			*r = Error
		}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_generalized_time_includes_fraction_seconds",
		Description:   "Generalized time values must not include fraction seconds",
		Providence:    "RFC 5280: 4.1.2.5.2",
		EffectiveDate: util.RFC2459Date,
		Test:          &generalizedTimeFraction{},
		updateReport: func(report *LintReport, result ResultStruct) {
			report.EGeneralizedTimeIncludesFractionSeconds = result
		},
	})
}
