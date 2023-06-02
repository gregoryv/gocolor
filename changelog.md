# Changelog
All notable changes to this project will be documented in this file.

The format is based on http://keepachangelog.com/en/1.0.0/
and this project adheres to http://semver.org/spec/v2.0.0.html.

## [unreleased]

- Colorize go doc output if firstline starts with "package "

## [0.7.0] - 2022-11-27

- COLOREXP supports multiple attributes, eg. .*:red;bgyellow
- Update help

## [0.6.0] - 2022-11-27

- Replace --custom option with optional named arguments

## [0.5.0] - 2022-11-27

- Add option --custom for coloring by regexp
- Use gregoryv/cmdline

## [0.4.0] - 2022-11-05

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
