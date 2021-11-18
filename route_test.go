package main

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestInitRoute(t *testing.T) {
	tests := []struct {
		name string
		want *gin.Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitRoute(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
