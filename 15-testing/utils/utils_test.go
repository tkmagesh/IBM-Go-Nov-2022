package utils

import "testing"

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func Test_isPrime_7(t *testing.T) {
	//arrange
	data := 7
	expectedResult := true

	//act
	actualResult := isPrime(data)

	//assert
	if actualResult != expectedResult {
		t.Errorf("Expected [%t] but received [%t]", expectedResult, actualResult)
	}
}

func Test_isPrime_10(t *testing.T) {
	//arrange
	data := 10
	expectedResult := false

	//act
	actualResult := isPrime(data)

	//assert
	if actualResult != expectedResult {
		t.Errorf("Expected [%t] but received [%t]", expectedResult, actualResult)
	}
}

func Test_isPrime(t *testing.T) {
	test_data := []struct {
		name     string
		data     int
		expected bool
		actual   bool
	}{
		{name: "is 9 prime", data: 9, expected: false},
		{name: "is 11 prime", data: 11, expected: true},
		{name: "is 13 prime", data: 13, expected: true},
		{name: "is 15 prime", data: 15, expected: false},
		{name: "is 17 prime", data: 17, expected: false},
		{name: "is 19 prime", data: 19, expected: true},
	}

	for _, td := range test_data {
		t.Run(td.name, func(t *testing.T) {
			td.actual = isPrime(td.data)

			//assert
			if td.actual != td.expected {
				t.Errorf("Expected [%t] but received [%t]", td.expected, td.actual)
			}
		})
	}
}
