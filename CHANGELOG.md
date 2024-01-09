# Changelog

## [0.1.3](https://github.com/cybroslabs/gpg-encryptor/compare/v0.1.2...v0.1.3) (2024-01-09)


### Bug Fixes

* **helm:** do not add whitespace before image tag ([440fe31](https://github.com/cybroslabs/gpg-encryptor/commit/440fe312f23d77594421180883556b924827527b))

## v0.1.2

### Bugfix

- fix(helm): properly handle helm appVersion, docker image has v prefix
- fix(helm): name of .chart template block
- fix(helm): reset tag to ""
- ci: remove unfinished workflows (release, publish-helm)

## v0.1.1

### Bugfix

- ci: build docker image on tags as well

## v0.1.0

First release - publish the docker image and helm chart.
