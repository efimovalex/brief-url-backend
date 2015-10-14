// +build acceptance

/*
	Each set of acceptance tests should be in their own file
	This file is a good place to put common test files, such as
	functions and structs that are re-used.

	There should be no actual tests here.
*/

package acceptance

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/efimovalex/brief_url/app"
)

var version string
var acceptanceConfig = &app.Config{}

const (
	SomeExampleConstMyAppTestsNeed = "ffae4c32-c90d-11e4-85d5-0800278bd2f6"
)

func init() {
	// Read in the version; this allows us to verify the we are testing the correct binary
	b, err := ioutil.ReadFile("../version")
	if err != nil {
		fmt.Printf("*** unexpected err reading version file - %s:%d ***", err)
	}
	version = strings.TrimSpace(string(b))

	// Load up the required acceptance environment variables
	envy.LoadWithPrefix("BRIEF_URL_", acceptanceConfig)
}

func SomeHelperFunc() {}
