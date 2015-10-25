package app

// This is the version for the package.
// some build tricks in ./bin/build will populate
// these variables allowing the healthcheck endpoint
// to show the compiled version and build number

var VERSION string = "VERSION_NOT_DEFINED"

var BUILD_NUMBER string = "BUILD_NUMBER_NOT_DEFINED"
