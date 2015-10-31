BriefURL
==========================


usage on Dev Environment
==========================

####Basic Usage

Using bin scripts to build and test
- `./bin/build`
- `./bin/test`
- `./bin/test_acceptance` # depends on other running services and potentially network services

You can start upstream services with `./bin/init_acceptance`

And stop upstream services with `./bin/destroy_acceptance`


####To run my_app:

After running ```./bin/build``` you can start this service with:

``` ./build/brief_url ```

usage on Local
==========================

Using GoDep with the following commnands:
- `godep go build`
- `godep go test` # optional build tags -tags=acceptance
or
- `godep restore; go test ./...`

Note: After you have built the binary, it will run on localhost:50111
- run with `./brief_url`

