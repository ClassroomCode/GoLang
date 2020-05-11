package format_test

import (
	"fmt"

	format "github.com/gopherguides/learn/_training/testing/example/src/solution"
)

// section: exercise
func ExampleTemplate() {
	t := `%key: %value`
	m := map[string]interface{}{"key": "Name", "value": "Rob Pike"}
	t = format.Template(t, m)
	fmt.Printf(t)
	// Output: Name: Rob Pike

}

/*
 TODO: Create another example that will show up as `Template (Email)` in the documentation
 Use the following template: `%name<%email>`

 Use the following values for the map:
  name: Rob Pike
  email: commander@google.com

 Expect it to have the following output: Rob Pike<commander@google.com>
*/

/*
 TODO: Create another example this time with the expected name of `Template (Html)`

 Use the following template:

`<html>
<body>
<h1>Hello %name</h1>
Thank you for subscribing to our email.

You can always unsubscribe here in the future:
<href a="%unsubscribe_link">Unsubscribe</a>
</body>
</html> `

Use the following key/values:
		"name":             "Rob Pike"
		"unsubscribe_link": "https://www.gopherguides.com/unsubscribe?email=commander@google.com"

Expect it to have the following output:

<html>
<body>
<h1>Hello Rob Pike</h1>
Thank you for subscribing to our email.

You can always unsubscribe here in the future:
<href a="https://www.gopherguides.com/unsubscribe?email=commander@google.com">Unsubscribe</a>
</body>
</html>

*/

// section: exercise
