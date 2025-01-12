# go-ipfs-example-plugin

[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](https://protocol.ai)
[![](https://img.shields.io/badge/project-IPFS-blue.svg?style=flat-square)](https://ipfs.io/)
[![](https://img.shields.io/badge/freenode-%23ipfs-blue.svg?style=flat-square)](http://webchat.freenode.net/?channels=%23ipfs)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![GoDoc](https://godoc.org/github.com/ipfs/go-ipfs-example-plugin?status.svg)](https://godoc.org/github.com/ipfs/go-ipfs-example-plugin)

> example plugin for Kubo

This repository contains a set of example Kubo plugins, demonstrating how to integrate external plugins into Kubo.

Packages:

* delaystore: an example datastore plugin that wraps an inner datastore in a "delayed" datastore.
* greeter: an example daemon plugin that prints "Hello!" on start and "Goodbye!" on exit.

**NOTE 1:** Plugins only work on Linux and MacOS at the moment. You can track the progress of this issue here: https://github.com/golang/go/issues/19282

**NOTE 2:** This plugin exists as an *example* and a starting point for new plugins. It isn't particularly useful by itself.

## Building and Installing

You can *build* the example plugin by running `make build`. You can then install it into your local IPFS repo by running `make install`.

Plugins need to be built against the correct version of Kubo. This package generally tracks the latest Kubo release but if you need to build against a different version, please set the `IPFS_VERSION` environment variable.

You can set `IPFS_VERSION` to:

* `vX.Y.Z` to build against that version of IPFS.
* `$commit` or `$branch` to build against a specific Kubo commit or branch.
   * Note: if building against a commit or branch make sure to build that commit/branch using the -trimpath flag. For example getting the binary via `go get -trimpath github.com/ipfs/kubo/cmd/ipfs@COMMIT`
* `/absolute/path/to/source` to build against a specific Kubo checkout.

To update the Kubo, run:

```bash
> make go.mod IPFS_VERSION=version
```

## Contribute

Feel free to join in. All welcome. Open an [issue](https://github.com/ipfs/go-ipfs-example-plugin/issues)!

This repository falls under the IPFS [Code of Conduct](https://github.com/ipfs/community/blob/master/code-of-conduct.md).

### Want to hack on IPFS?

[![](https://cdn.rawgit.com/jbenet/contribute-ipfs-gif/master/img/contribute.gif)](https://github.com/ipfs/community/blob/master/CONTRIBUTING.md)

## License

MIT
