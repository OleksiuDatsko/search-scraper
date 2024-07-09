package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDomainFRomURL(t *testing.T) {
	type args struct {
		url string
	}

	tests := map[string]struct {
		args
		want string
	}{
		"valid_url": {
			args: args{
				url: "http://example.com",
			},
			want: "example.com",
		},
		"invalid_url": {
			args: args{
				url: "example.com",
			},
			want: "example.com",
		},
		"empty_url": {
			args: args{
				url: "",
			},
			want: "",
		},
		"no_url": {
			args: args{
				url: "",
			},
			want: "",
		},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {
			got := GetDomainFromURL(tt.args.url)
			assert.Equal(t, tt.want, got)
		})
	}
}
