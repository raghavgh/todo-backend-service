package server

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

// Unit test of BuildRouter function
func TestBuildRouter(t *testing.T) {
	tests := []struct {
		name string
		want *mux.Router
	}{
		{
			name: "TestBuildRouter",
			want: mux.NewRouter(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewServer(t *testing.T) {
	type args struct {
		router *mux.Router
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			name: "TestNewServer",
			args: args{
				router: mux.NewRouter(),
			},
			want: &Server{
				Addr:         ":8080",
				Handler:      mux.NewRouter(),
				ReadTimeout:  1000000000,
				WriteTimeout: 1000000000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.router); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStart(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Start()
		})
	}
}

func Test_enableCors(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := enableCors(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("enableCors() = %v, want %v", got, tt.want)
			}
		})
	}
}
