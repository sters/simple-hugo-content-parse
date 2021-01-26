package hugocontent

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	yaml "github.com/goccy/go-yaml"
	"github.com/morikuni/failure"
)

// MarkdownContent for hugo
// https://gohugo.io/content-management/formats/
type MarkdownContent struct {
	// FrontMatter is metadata for this content
	FrontMatter map[string]interface{}
	// Body for this content
	Body string
}

// Dump to string from this content
func (m *MarkdownContent) Dump() (string, error) {
	meta, err := yaml.Marshal(m.FrontMatter)
	if err != nil {
		return "", failure.Wrap(err, failure.WithCode(ErrFileContentMismatch))
	}

	return fmt.Sprintf(`%s%s
%s%s`,
		hugoSeparator, strings.TrimSpace(string(meta)), hugoSeparator, m.Body), nil
}

// see https://gohugo.io/content-management/front-matter/#front-matter-formats
const hugoSeparator = "---\n"

var (
	// ErrFileContentMismatch on specified filepath
	ErrFileContentMismatch = failure.StringCode("file content mismatch")
)

// ParseMarkdownWithYaml from any reader to make MarkdownContent struct
func ParseMarkdownWithYaml(r io.Reader) (*MarkdownContent, error) {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	content := strings.Split(string(raw), hugoSeparator)
	if len(content) < 3 {
		return nil, failure.New(ErrFileContentMismatch)
	}

	c := &MarkdownContent{
		Body: strings.Join(content[2:], hugoSeparator),
	}

	if err := yaml.Unmarshal([]byte(content[1]), &c.FrontMatter); err != nil {
		return nil, failure.Wrap(err, failure.WithCode(ErrFileContentMismatch))
	}

	return c, nil
}
