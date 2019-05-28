package main

import (
	"testing"
)

func Test_office_getDistance(t *testing.T) {
	o := office{
		Location: location{-90.000000, 0.000000},
	}
	type args struct {
		start location
		end   location
	}
	tests := []struct {
		name   string
		fields office
		args   location
		want   float64
	}{
		{
			name:   "zero is distance between same start and end location",
			fields: o,
			args:   location{-90.000000, 0.000000},
			want:   0,
		},
		{
			name:   "distance north pole and south pole",
			fields: o,
			args:   location{90.000000, 0.000000},
			want:   20015.086796020572,
		},
		{
			name:   "distance dublin and london",
			fields: office{Location: location{53.33306, -6.24889}},
			args:   location{51.50853, -0.12574},
			want:   15735.030612929339,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := o.getDistance(tt.args); got != tt.want {
				t.Errorf("office.getDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
