package Ajit

import (
	"testing"
)

func TestEmployeeSalary(t *testing.T) {
	tests := []struct {
		name string
		//want  *os.File
		want1 bool
	}{
		// TODO: Add test cases.
		{
			name:  "testing the file size",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := EmployeeSalary()

			if got1 != tt.want1 {
				t.Errorf("EmployeeSalary() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
