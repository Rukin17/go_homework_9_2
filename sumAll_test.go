package main

import "testing"

func TestSumAll(t *testing.T) {
	var testCases = []struct {
		name       string
		nums       []int
		wantResult int
	}{
		{"positive numbers", []int{1, 2, 3}, 6},
		{"negative numbers", []int{-1, -2, -3}, -6},
		{"empty slice", []int{}, 0},
		{"single number", []int{17}, 17},
		{"positive & negative numbers", []int{2, -3}, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := sumAll(tc.nums...)
			if res != tc.wantResult {
				t.Errorf("Ожидался результат %v, получен : %v", tc.wantResult, res)
			}
		})
	}
}

func TestConvertStringToSlice(t *testing.T) {
	var testCases = []struct {
		name       string
		input      string
		wantResult []int
		wantErr    bool
	}{
		{"positive numbers", "1 2 3", []int{1, 2, 3}, false},
		{"negative numbers", "-1 -2 -3", []int{-1, -2, -3}, false},
		{"empty slice", "", []int{}, false},
		{"single number", "17", []int{17}, false},
		{"positive & negative numbers", "2 -3", []int{2, -3}, false},
		{"invalid numbers", "22 f4", []int{22, 4}, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := convertStringToSlice(tc.input)
			for i := range res {
				if res[i] != tc.wantResult[i] {
					t.Errorf("Ожидалоcь %v, получено: %v", tc.wantResult, res)
				}
			}
		})
	}
}

func TetsInput(t *testing.T) {
	// Хотел замокать input(), но процент покрытия тестами не изменился..
	// Так и должно быть или я что-то не так сделал?
	t.Skip("Ввод из консоли замокан")
}
