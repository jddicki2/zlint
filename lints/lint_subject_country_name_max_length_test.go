// lint_serial_number_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectCountryNameLengthGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectCountryNameLengthGood.pem"
	desEnum := Pass
	out, _ := Lints["e_subject_country_name_max_length"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubjectCountryNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectCountryNameLong.pem"
	desEnum := Error
	out, _ := Lints["e_subject_country_name_max_length"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
