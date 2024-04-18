package myutils

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestCheckValidDomain(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"TestCheckValidDomain", args{"http://www.google.com"}, true},
		{"TestCheckValidDomain", args{"https://www.google.com"}, true},
		{"TestCheckValidDomain", args{"http://www.google.com:8080"}, true},
		{"TestCheckValidDomain", args{"http://www.google.com:8080/abc"}, true},
		{"TestCheckValidDomain", args{"http://www.google.com:8080/abc?def=123"}, true},
		{"TestCheckValidDomain", args{"https://www.google.com:8080/abc?def=123#456"}, true},
		{"TestCheckValidDomain", args{"http://www.google.com:8080/abc?def=123#456/xyz/abc"}, true},
		{"TestCheckValidDomain", args{"www.google.com"}, false},
		{"TestCheckValidDomain", args{"com"}, false},
		{"TestCheckValidDomain", args{""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckValidDomain(tt.args.domain); got != tt.want {
				t.Errorf("CheckValidDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckValidIp(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"TestCheckValidIp", args{"222.222.222.222"}, true},
		{"TestCheckValidIp", args{"1.1.1.1.1"}, false},
		{"TestCheckValidIp", args{"1.1.1"}, false},
		{"TestCheckValidIp", args{"1.1.1."}, false},
		{"TestCheckValidIp", args{"1.1.1.256"}, false},
		{"TestCheckValidIp", args{"255.255.255.255"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckValidIp(tt.args.ip); got != tt.want {
				t.Errorf("CheckValidIp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTopLevelDomain(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestGetTopLevelDomain", args{"http://www.google.com"}, "com", false},
		{"TestGetTopLevelDomain", args{"https://www.google.com"}, "com", false},
		{"TestGetTopLevelDomain", args{"http://www.google.com:8080"}, "com", false},
		{"TestGetTopLevelDomain", args{"http://www.google.com:8080/abc"}, "com", false},
		{"TestGetTopLevelDomain", args{"http://www.google.com:8080/abc?def=123"}, "com", false},
		{"TestGetTopLevelDomain", args{"https://www.google.com:8080/abc?def=123#456"}, "com", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTopLevelDomain(tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopLevelDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetTopLevelDomain() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildResponse(t *testing.T) {
	type args struct {
		status  bool
		message string
		data    map[string]string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "TestBuildResponse success",
			args: args{
				status:  true,
				message: "Success",
				data:    map[string]string{"key": "value"},
			},
			want: func() []byte {
				response := Response{
					Code: Success,
					Msg:  "Success",
					Data: map[string]string{"key": "value"},
				}
				responseJson, _ := json.Marshal(response)
				return responseJson
			}(),
		},
		{
			name: "TestBuildResponse failed",
			args: args{
				status:  false,
				message: "Failed",
				data:    map[string]string{"key": "value"},
			},
			want: func() []byte {
				response := Response{
					Code: Failed,
					Msg:  "Failed",
					Data: map[string]string{"key": "value"},
				}
				responseJson, _ := json.Marshal(response)
				return responseJson
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildResponse(tt.args.status, tt.args.message, tt.args.data); !bytes.Equal(got, tt.want) {
				t.Errorf("BuildResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
