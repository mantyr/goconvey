package assertions

import (
	"fmt"
	"reflect"
	"strings"
)

// ShouldStartWith receives exactly 2 string parameters and ensures that the first starts with the second.
func ShouldStartWith(actual interface{}, expected ...interface{}) string {
	if fail := need(1, expected); fail != success {
		return fail
	}

	value, valueIsString := actual.(string)
	prefix, prefixIsString := expected[0].(string)

	if !valueIsString || !prefixIsString {
		return fmt.Sprintf(shouldBothBeStrings, reflect.TypeOf(actual), reflect.TypeOf(expected[0]))
	}

	return shouldStartWith(value, prefix)
}
func shouldStartWith(value, prefix string) string {
	if !strings.HasPrefix(value, prefix) {
		return fmt.Sprintf(shouldHaveStartedWith, value, prefix)
	}
	return success
}

// ShouldNotStartWith receives exactly 2 string parameters and ensures that the first does not start with the second.
func ShouldNotStartWith(actual interface{}, expected ...interface{}) string {
	if fail := need(1, expected); fail != success {
		return fail
	}

	value, valueIsString := actual.(string)
	prefix, prefixIsString := expected[0].(string)

	if !valueIsString || !prefixIsString {
		return fmt.Sprintf(shouldBothBeStrings, reflect.TypeOf(actual), reflect.TypeOf(expected[0]))
	}

	return shouldNotStartWith(value, prefix)
}
func shouldNotStartWith(value, prefix string) string {
	if strings.HasPrefix(value, prefix) {
		if value == "" {
			value = "<empty>"
		}
		if prefix == "" {
			prefix = "<empty>"
		}
		return fmt.Sprintf(shouldNotHaveStartedWith, value, prefix)
	}
	return success
}

// ShouldEndWith receives exactly 2 string parameters and ensures that the first ends with the second.
func ShouldEndWith(actual interface{}, expected ...interface{}) string {
	if fail := need(1, expected); fail != success {
		return fail
	}

	value, valueIsString := actual.(string)
	suffix, suffixIsString := expected[0].(string)

	if !valueIsString || !suffixIsString {
		return fmt.Sprintf(shouldBothBeStrings, reflect.TypeOf(actual), reflect.TypeOf(expected[0]))
	}

	return shouldEndWith(value, suffix)
}
func shouldEndWith(value, suffix string) string {
	if !strings.HasSuffix(value, suffix) {
		return fmt.Sprintf(shouldHaveEndedWith, value, suffix)
	}
	return success
}

// ShouldEndWith receives exactly 2 string parameters and ensures that the first does not end with the second.
func ShouldNotEndWith(actual interface{}, expected ...interface{}) string {
	if fail := need(1, expected); fail != success {
		return fail
	}

	value, valueIsString := actual.(string)
	suffix, suffixIsString := expected[0].(string)

	if !valueIsString || !suffixIsString {
		return fmt.Sprintf(shouldBothBeStrings, reflect.TypeOf(actual), reflect.TypeOf(expected[0]))
	}

	return shouldNotEndWith(value, suffix)
}
func shouldNotEndWith(value, suffix string) string {
	if strings.HasSuffix(value, suffix) {
		if value == "" {
			value = "<empty>"
		}
		if suffix == "" {
			suffix = "<empty>"
		}
		return fmt.Sprintf(shouldNotHaveEndedWith, value, suffix)
	}
	return success
}