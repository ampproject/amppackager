# Contributing Guidelines

We'd love to accept your patches and contributions to this project. There are
just a few small guidelines you need to follow.

## Contributor License Agreement

Contributions to this project must be accompanied by a Contributor License
Agreement. You (or your employer) retain the copyright to your contribution;
this simply gives us permission to use and redistribute your contributions as
part of the project. Head over to <https://cla.developers.google.com/> to see
your current agreements on file or to sign a new one.

You generally only need to submit a CLA once, so if you've already submitted one
(even if it was for a different project), you probably don't need to do it
again.

## Code reviews

All submissions, including submissions by project members, require review. We
use GitHub pull requests for this purpose. Consult
[GitHub Help](https://help.github.com/articles/about-pull-requests/) for more
information on using pull requests.

## Community Guidelines

This project follows [Google's Open Source Community
Guidelines](https://opensource.google.com/conduct/).

# Getting Started

## Running a development server

Copy `amppkg.example.toml` to `amppkg.toml` and modify it to suit your needs.
You may use `testdata/b1/server.{cert,privkey}` to get started. However, you are
at your own risk if you instruct your browser to trust `server.cert`. *NEVER*
instruct your browser to trust `ca.cert`.

## Presubmits

- Run `go fmt` on the code you change. Don't run `go fmt ./...`; this affects
  files that are mirrored from Google to GitHub, so we can't change them here.
- Make sure `go test ./...` passes.
- `golint` and `go vet` are optional.

## New dependencies

Feel free to add dependencies on small code if it implements a feature that's
hard to implement yourself. Try not to add large dependencies, or dependencies
for the sake of minor development inconvenience (unless it's for test code
only). They add risk by bringing in code of unknown provenance, and bloat the
binary.

If you need to add or upgrade dependencies, AMP Packager uses
[dep](https://golang.github.io/dep/).

# Suggestions for contributions

Take a look at [good first
issues](https://github.com/ampproject/amppackager/labels/good%20first%20issue),
and please communicate early and often, to ensure we agree on the solution
before you invest a lot of time into its implementation.
