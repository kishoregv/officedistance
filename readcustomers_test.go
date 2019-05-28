package main

import (
	"reflect"
	"testing"
)

func Test_readCustomersList(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []*customer
		wantErr bool
	}{
		{
			name:    "Read a non existing file Dummy.txt",
			args:    args{filename: "Dummy.txt"},
			want:    nil,
			wantErr: true,
		}, {
			name:    "Read an empty file empty.txt",
			args:    args{filename: "test-fixtures/empty.txt"},
			want:    nil,
			wantErr: true,
		}, {
			name:    "Read an invalid file",
			args:    args{filename: "test-fixtures/invalid_customer.txt"},
			want:    nil,
			wantErr: true,
		}, {
			name: "Read test fixture 2_customers.txt",
			args: args{filename: "test-fixtures/2_customers.txt"},
			want: []*customer{
				&customer{
					Lat: "52.986375", UserID: 12, Name: "Christina McArdle", Long: "-6.043701",
					location: location{lat: 52.986375, long: -6.043701},
				},
				&customer{
					Lat: "51.92893", UserID: 1, Name: "Alice Cahill", Long: "-10.27699",
					location: location{lat: 51.92893, long: -10.27699},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readCustomersList(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCustomersList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCustomersList() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
