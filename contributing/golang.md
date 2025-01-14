# Nomad Go Version

Nomad is intended to be built with a specific version of the Go toolchain for
each release. Generally, each Y release of Nomad (where 0.9.5 means X=0, Y=9,
Z=5) will update to the latest version of the Go toolchain available at the
time.

Nomad Z releases update to the latest Go Z release but do *not* change Go's Y
version.

## Version Table

| Nomad Version | Go Version |
|:-------------:|:----------:|
| 1.2           | 1.17       |
| 1.1           | 1.16       |
| 1.0           | 1.15       |
| 0.12          | 1.14       |
| 0.11          | 1.14       |
| 0.10          | 1.12       |
| 0.9           | 1.11       |

## Code

The
[`update_golang_version.sh`](https://github.com/hernad/nomad/blob/main/scripts/update_golang_version.sh)
script is used to update the Go version for all build tools.

The [Changelog](https://github.com/hernad/nomad/blob/main/CHANGELOG.md)
will note when the Go version has changed in the Improvements section:

```
* build: Updated to Go 1.12.13 [GH-6606]
```
