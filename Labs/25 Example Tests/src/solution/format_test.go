package format_test

import (
	"fmt"

	format "github.com/gopherguides/learn/_training/testing/example/src/solution"
)

func ExampleTemplate() {
	t := `%key: %value`
	m := map[string]interface{}{"key": "Name", "value": "Rob Pike"}
	t = format.Template(t, m)
	fmt.Printf(t)
	// Output: Name: Rob Pike

}

func ExampleTemplate_email() {
	t := `%name<%email>`
	m := map[string]interface{}{"name": "Rob Pike", "email": "commander@google.com"}
	t = format.Template(t, m)
	fmt.Printf(t)
	// Output: Rob Pike<commander@google.com>

}

func ExampleTemplate_html() {
	t := `<html>
<body>
<h1>Hello %name</h1>
Thank you for subscribing to our email.

You can always unsubscribe here in the future:
<href a="%unsubscribe_link">Unsubscribe</a>
</body>
</html> `
	m := map[string]interface{}{
		"name":             "Rob Pike",
		"unsubscribe_link": "https://www.gopherguides.com/unsubscribe?email=commander@google.com",
	}
	t = format.Template(t, m)

	fmt.Printf(t)

	// Output: <html>
	// <body>
	// <h1>Hello Rob Pike</h1>
	// Thank you for subscribing to our email.
	//
	// You can always unsubscribe here in the future:
	// <href a="https://www.gopherguides.com/unsubscribe?email=commander@google.com">Unsubscribe</a>
	// </body>
	// </html>

}
