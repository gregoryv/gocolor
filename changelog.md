# Changelog
All notable changes to this project will be documented in this file.

The format is based on http://keepachangelog.com/en/1.0.0/
and this project adheres to http://semver.org/spec/v2.0.0.html.

## [unreleased]

- Color coverage statement result, green if 100%, magenta otherwise
- Add usage to cmd/gocolor

## [0.3.0] - 2022-03-24

- Optimize func Colorize, 1 allocs/op

## [0.2.0] - 2022-03-23

- Color subtest variations
- Keep the ':' as is, but color it

## [0.1.1] - 2022-03-21

- Remove colon after --- FAIL:

## [0.1.0] - 2022-03-21

- Use cyan for skipped tests
- Use exit code 1 if test failure is found
- Add cmd/gocolor
- Use yellow, green and red colors
