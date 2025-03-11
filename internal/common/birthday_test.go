package common_test

import (
	"ingressos/internal/common"
	"testing"
	"time"
)

type testCase struct {
	date        time.Time
	now         time.Time
	expectedAge int
}
type testCases = []testCase

func TestBirtday(t *testing.T) {
	cases := testCases{
		{date: time.Date(2003, 8, 26, 0, 0, 0, 0, time.UTC), now: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), expectedAge: 21},
		{date: time.Date(2004, 8, 26, 0, 0, 0, 0, time.UTC), now: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), expectedAge: 20},
	}

	for _, testCase := range cases {
		birthday, err := common.CreateBirthday(testCase.date, common.FakeClock{MockedNow: testCase.now})
		if err != nil {
			t.Fatal("unexpected error")
		}

		age := birthday.GetAge()
		if age != testCase.expectedAge {
			t.Fatalf("expected age of %d, but received %d", testCase.expectedAge, age)
		}
	}
}
