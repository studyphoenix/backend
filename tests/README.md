# Tests
To run tests, the backend needs to be running already. For CI, a backend will be spun up to test against.
This is only run on Pull Requests into main because it takes quite a while to run.

## Components
This tests the components individually, without assignments

## E2E
This executes end to end tests, including associations and global deletes (i.e. checks user deletion, etc)

TODO: Access Control testing!!

## CI
Testing via CI will be added at some point in the near future

## Writing tests
All tests use the OpenCollection format, see the [bruno docs](https://docs.usebruno.com/opencollection-yaml/overview)
and the official [OpenCollection Spec](https://spec.opencollection.com/#introduction)
