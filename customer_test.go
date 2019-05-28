package main

import (
	"reflect"
	"testing"
)

func Test_customer_unmarshal(t *testing.T) {

	tests := []struct {
		name    string
		args    []byte
		want    *customer
		wantErr bool
	}{
		{
			name:    "Empty customer",
			args:    []byte{},
			want:    &customer{},
			wantErr: true,
		},
		{
			name:    "invalid customer line",
			args:    []byte("{"),
			want:    &customer{},
			wantErr: true,
		},
		{
			name:    "invalid latitute in customer line",
			args:    []byte("{\"latitude\": \"xyz\", \"user_id\": 1, \"name\": \"C1\", \"longitude\": \"0.0\"}"),
			want:    &customer{Lat: "xyz", Long: "0.0", UserID: 1, Name: "C1", location: location{}},
			wantErr: true,
		},
		{
			name:    "Basic customer",
			args:    []byte("{\"latitude\": \"90.0\", \"user_id\": 1, \"name\": \"C1\", \"longitude\": \"0.0\"}"),
			want:    &customer{Lat: "90.0", Long: "0.0", UserID: 1, Name: "C1", location: location{90.0, 0.0}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &customer{}
			if err := c.unmarshal(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("customer.unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("readCustomersList() = %+v, want %+v", c, tt.want)
			}
		})
	}
}
