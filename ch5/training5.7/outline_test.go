package training5_7

import "testing"

func TestOutline(t *testing.T) {
	dom := `<html><head><meta charset="utf8" /><title>test</title><script type="text/javascript" src="test.js"></script></head><body><h1>test</h1><div><p><a href="test.html">test</a></p><img src="test.jpg" /></div></body></html>`
	expected := `<html>
  <head>
    <meta charset="utf8"/>
    <title>
    </title>
    <script type="text/javascript" src="test.js">
    </script>
  </head>
  <body>
    <h1>
    </h1>
    <div>
      <p>
        <a href="test.html">
        </a>
      </p>
      <img src="test.jpg"/>
    </div>
  </body>
</html>
`

	actual := outline(dom)

	if expected != actual {
		t.Errorf("expected:%#v\n\nactual:%#v", expected, actual)
	}
}