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
		{"t5", `
 <body class="nodata company_blog" style="">
        <script>            var toolbarSearchExt = '{"landingWord":[],"queryWord":"","tag":["function","class","filter","search"],"title":"Yii: 设置数据翻页"}';
        </script>
    <script src="https://g.csdnimg.cn/common/csdn-toolbar/csdn-toolbar.js" type="text/javascript"></script>
    <script src="https://g.csdnimg.cn/common/csdn-toolbar/csdn-toolbar1.js" type="text/javascript"></script>
    <script>
    (function(){
        var bp = document.createElement('script');
        var curProtocol = window.location.protocol.split(':')[0];
        if (curProtocol === 'https') {
            bp.src = 'https://zz.bdstatic.com/linksubmit/push.js';
        }
        else {
            bp.src = 'http://push.zhanzhang.baidu.com/push.js';
        }
        var s = document.getElementsByTagName("script")[0];
        s.parentNode.insertBefore(bp, s);
    })();
    </script>
<link rel="stylesheet" href="https://csdnimg.cn/release/blogv2/dist/pc/css/blog_code-01256533b5.min.css">
<link rel="stylesheet" href="https://csdnimg.cn/release/blogv2/dist/mdeditor/css/editerView/chart-3456820cac.css" /><div style='font-size: 12px;'>hello</div></body>`, "hello"},
		{"t6", `<!-- show up to 2 reviews by default -->












<p>Custom flags for your garden are a great way to show your personality to your friends and neighbors. Design and turn it into an eye-catching flag all year round. This will be a beautiful addition to your yard and garden, also a simple sign to show your patriotism on Memorial Day, 4th of July or Veterans Day, Christmas holidays or any holiday of the year.

</p>`, "Custom flags for your garden are a great way to show your personality to your friends and neighbors. Design and turn it into an eye-catching flag all year round. This will be a beautiful addition to your yard and garden, also a simple sign to show your patriotism on Memorial Day, 4th of July or Veterans Day, Christmas holidays or any holiday of the year."},
		{"t7", "&lt;div>hello<div>", "hello"},
	}

	for _, test := range tests {
		equal := Strip(test.html)
		assert.Equal(t, test.expected, equal, test.tag)
	}
}

func BenchmarkStrip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Strip(`<!-- show up to 2 reviews by default -->












<p>Custom flags for your garden are a great way to show your personality to your friends and neighbors. Design and turn it into an eye-catching flag all year round. This will be a beautiful addition to your yard and garden, also a simple sign to show your patriotism on Memorial Day, 4th of July or Veterans Day, Christmas holidays or any holiday of the year.

</p>`)
	}
}

func TestTag(t *testing.T) {
	tests := []struct {
		tag        string
		elementTag string
		content    string
		attributes map[string]string
		styles     map[string]string
		expected   string
	}{
		{"t0", "div", "hello", nil, nil, "<div>hello</div>"},
		{"t1", "div", "hello", map[string]string{"id": "name"}, nil, `<div id="name">hello</div>`},
		{"t1.1", "div", "hello", map[string]string{"id": "name", "name": "name"}, nil, `<div id="name" name="name">hello</div>`},
		{"t2", "div", "hello", map[string]string{"id": "name", "data-tag": "123"}, map[string]string{"font-size": "1", "font-weight": "bold"}, `<div data-tag="123" id="name" style="font-size:1;font-weight:bold;">hello</div>`},
	}

	for _, test := range tests {
		equal := Tag(test.elementTag, test.content, test.attributes, test.styles)
		assert.Equal(t, test.expected, equal, test.tag)
	}
}

func BenchmarkTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tag("div", "hello", map[string]string{"id": "name"}, map[string]string{"font-size": "1"})
	}
}
