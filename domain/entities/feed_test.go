package entities

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewFeed(t *testing.T) {
	type args struct {
		title       string
		description string
		link        string
		date        string
	}
	tests := []struct {
		name string
		args args
		want *Feed
	}{
		{
			name: "Should create a new Feed",
			args: args{
				title:       "Test Title",
				description: "Test Description",
				link:        "https://catracalivre.com.br/feed/",
				date:        "Thu, 12 Jan 2023 00:19:41 +0000",
			},
			want: &Feed{
				Title:       "Test Title",
				Description: "Test Description",
				Provider:    "catracalivre",
				Link:        "https://catracalivre.com.br/feed/",
				Date:        "Thu, 12 Jan 2023 00:19:41 +0000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFeed(tt.args.title, tt.args.description, tt.args.link, tt.args.date)
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(Feed{}, "CreatedAt")); diff != "" {
				t.Errorf("%v mismatch (-want +got):\n%s", tt.name, diff)
			}
		})
	}
}
