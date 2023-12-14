Easily convert `spdx.json` files into human-readable, markdown files.

## Installation

```shell
go install go.pitz.tech/spdx-fmt@latest
```

### Usage

Usage of this tool assumes the usage of one of the many other SPDX tools used to generate JSON documents. Using those
files, we can generate human-readable versions of their contents as well as artifacts useful for shipping software (such
as reporting on third-party licenses).

```shell
Usage of spdx-fmt:
  -input string
        Specify the input spdx.json file. Defaults to '-' for stdin. (default "-")
  -output string
        Specify the output markdown file. Defaults to '-' for stdout. (default "-")
  -report string
        Specify which report to render (spdx, third-party-licenses, or a file path). (default "spdx")
```
