package hugocontent

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParseMarkdownWithYaml(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    *MarkdownContent
		wantErr bool
	}{
		{
			name:    "empty",
			raw:     "",
			want:    nil,
			wantErr: true,
		},
		{
			name: "simple",
			raw: `---
foo: bar
baz: 1
---

foo
`,
			want: &MarkdownContent{
				FrontMatter: map[string]interface{}{
					"foo": "bar",
					"baz": uint64(1),
				},
				Body: `
foo
`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMarkdownWithYaml(bytes.NewBuffer([]byte(tt.raw)))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMarkdownWithYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("ParseMarkdownWithYaml() got = %+v, want %+v", got, tt.want)
				return
			}
		})
	}
}

func TestMarkdownContent_Dump(t *testing.T) {
	tests := []struct {
		name    string
		m       *MarkdownContent
		want    string
		wantErr bool
	}{
		{
			name: "simple",
			m: &MarkdownContent{
				FrontMatter: map[string]interface{}{
					"foo": "bar",
					"baz": 1,
				},
				Body: `
foo

bar
`,
			},
			want: `---
baz: 1
foo: bar
---

foo

bar
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Dump()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarkdownContent.Dump() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MarkdownContent.Dump() = %v, want %v", got, tt.want)
			}
		})
	}
}
