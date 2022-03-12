# AllTaxi backend application documentation

### Prerequisities:
- `go` with version above **1.14**
- `make` binary package
- `docker.io` (recommended)

To run the project, first you need to create `.env` file by running the following command:

    make create-env

Customize `.env` file according to your needs by replacing values for environment variables. Then run command:

    set -a &&. ./.env && set +a

To run the project in _development mode_:

    $ go run cmd/main.go

    or

    $ make run-dev

To generate protos:

    make proto-gen

To pull recent changes in protos submodule:

    make update-proto-module

To run SQL migrations locally:

    make migrate-jeyran
For running tests locally, linting your code, unit testing, race detection and memory sanitizing:

    make lint
    
    make unit-tests

    make race

    make msan

To delete unnecessary branches locally, run:

    make delete-branches

To update `protos` submodule run:

    make pull-proto-module && make update-proto-module

To generate routes for swagger run:

    make swag-gen

To run redis locally (you need docker that runs without sudo):

    make run-redis

To stop it:

    make stop-redis
