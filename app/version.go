package app

// This is the version for the package.
// some build tricks in ./bin/build will populate
// these variables allowing the healthcheck endpoint
// to show the compiled version and build number

var VERSION string = "VERSION NOT DEFINED"

var BUILD_NUMBER string = "BUILD NUMBER NOT DEFINED"
