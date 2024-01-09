# Changelog

## v0.1.3

### Bugfixes

- fix(helm): do not add whitespace before image tag

## v0.1.2

### Bugfixes

- fix(helm): properly handle helm appVersion, docker image has v prefix
- fix(helm): name of .chart template block
- fix(helm): reset tag to ""
- ci: remove unfinished workflows (release, publish-helm)

## v0.1.1

### Bugfixes

- ci: build docker image on tags as well

## v0.1.0

First release - publish the docker image and helm chart.
