# Changelog

## v0.1.5

### Bugfixes

- fix(swagger): base path to remove duplicate /v1/v1 from url in swagger ui

## v0.1.4

### Other changes

- ci(build): docker image with git tag tag
- docs: add release checklist
- docs(changelog): move changelog to docs/ dir
- chore: remove .releaserc
- ci: remove release workflow, doing it manually
- ci(release): run release workflow on manual trigger only
- chore(.releaserc): do not add [skip ci] suffix

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
