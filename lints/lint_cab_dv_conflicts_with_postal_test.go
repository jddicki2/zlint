// lint_cert_policy_conflicts_with_postal_test.go
package lints

import (
	"testing"
)

func TestCertPolicyNotConflictWithPostal(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	desEnum := Pass
	out, _ := Lints["e_cab_dv_conflicts_with_postal"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertPolicyConflictsWithPostal(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValWithPostal.pem"
	desEnum := Error
	out, _ := Lints["e_cab_dv_conflicts_with_postal"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
