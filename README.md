My App - A skeleton application for SendGrid API services
==========================

This is a skeleton application that may help inform how you set up your api service.

usage on Dev Environment
==========================

####Basic Usage

See Fortress Development for the most up-to-date dev environment.
Using bin scripts to build and test
- `./bin/build`
- `./bin/test`
- `./bin/test_acceptance` # depends on other running services and potentially network services

You can start upstream services with `./bin/init_acceptance`

And stop upstream services with `./bin/destroy_acceptance`


####To run my_app:

After running ```./bin/build``` you can start this service with:

``` $ . example.env && ./build/my_app ```

Healthcheck test by:

``` $ curl localhost:50112/healthcheck ```

usage on Local
==========================

Using GoDep with the following commnands:
- `godep go build`
- `godep go test` # optional build tags -tags=acceptance
or
- `godep restore; go test ./...`

Note: After you have built the binary, it will run on localhost:50111
- run with `./my_app`

Feature Toggles
=========================

To enable all Feature Toggles, run `./bin/enable_feature_toggles`. Any registered toggle will be turned on.


Testing and Feature Toggles
=========================

There are two location where you need to list your feature toggle.

First, in `acceptance/feature_toggle_test.go`, add the feature to the list of features used in `getFeatureToggleStates()`

Second, in `bin/enable_feature_toggles`, add it to the `arr` array.

After this, you are free to test your toggle. It is recommended that in your acceptance test, you do the following at the top of your test functions (example for "paid_signup" toggle):
```go
// TODO: re-add t.Parallel and remove toggle code after
// "paid_signup" is no longer a toggle.
// t.Parallel()
ts := getFeatureToggleStates()
defer restoreFeatureToggleStates(ts)
err := enableFeatureToggle("paid_signup")
if err != nil {
  t.Error("unable to set `paid_signup` toggle ", err)
}
```

And in your unit tests, it is recommended that you do the following faking for apid:
```go
fakeClient.RegisterFunction("checkFeatureToggle", func(params url.Values) (interface{}, error) {
    return true, nil
  })
```

Endpoints
==========================
- See the routes.go file
- Or see the [Apiary doc](apiary.link.here)

Contributing
==========================
See the [CONTRIBUTING.md](CONTRIBUTING.md) file
