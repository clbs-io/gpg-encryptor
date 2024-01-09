# Release

A release checklist.

## Format

- tag: `v\d+\.\d+\.\d+`
- commit (updating changelog, etc.) `chore(release): ${VERSION}` (`${VERSION}` has the same format as tag)

## Release process

- update changelog
- update helm chart version
- release commit
- tag that commit
- push
- create github release, copy the changelog for given release and add docker image URLs
