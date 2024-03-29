// Copyright 2015 Alex Browne and Soroush Pour.
// Allrights reserved. Use of this source code is
// governed by the MIT license, which can be found
// in the LICENSE file.

package vrouter_test

import (
	"fmt"
	"github.com/pubgo/vrouter"
)

func ExampleRouter_HandleFunc() {
	// Create a new Router object
	r := vrouter.New()
	// Use HandleFunc to add routes.
	r.HandleFunc("/greet/{name}", func(context *vrouter.Context) {
		// The handler for this route simply grabs the name parameter
		// from the map of params and says hello.
		fmt.Printf("Hello, %s\n", context.Params["name"])
	})
	// You must call Start in order to start listening for changes
	// in the url and trigger the appropriate handler function.
	r.Start()
}
