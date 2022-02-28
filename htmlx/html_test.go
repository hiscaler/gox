package htmlx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrip(t *testing.T) {
	tests := []struct {
		tag      string
		html     string
		expected string
	}{
		{"t0", "<div>hello</div>", "hello"},
		{"t1", `

<div>hello</div>

`, "hello"},
		{"t3", "<div style='font-size: 12px;'>hello</div>", "hello"},
		{"t4", "<style>body {font-size: 12px}</style><div style='font-size: 12px;'>hello</div>", "hello"},
		{"t4", `
<link rel='stylesheet' id='wp-block-library-css'  href='https://www.example.com/style.min.css?ver=5.9.1' type='text/css' media='all' />
<style type="text/css">body {font-size: 12px}</style><!-- / See later. --><div style='font-size: 12px;'>hello</div>`, "hello"},
	}

	for _, test := range tests {
		equal := Strip(test.html)
		assert.Equal(t, test.expected, equal, test.tag)
	}
}
