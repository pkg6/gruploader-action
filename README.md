
## GitHub action (file upload)

~~~
name: upload file extra

on:
  release:
    types: [created]

permissions:
    contents: write
    packages: write

jobs:
  release-linux-amd64:
    name: release linux/amd64
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
    - uses: pkg6/gruploader-action@main
      with:
        github_token: ${{ secrets.GH_TOKEN }}
        action: files
        extra_files: 
          README.md
          LICENSE
~~~

## GitHub action (golang upload)

~~~
name: build golang

on:
  release:
    types: [created]

jobs:
  build-go-binary:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        goos: [linux, windows]
        goarch: [arm64]
        exclude:
          - goarch: arm64
            goos: windows
    steps:
      - uses: actions/checkout@v3
      - uses: pkg6/gruploader-action@main
        with:
          github_token: ${{ secrets.GH_TOKEN }} 
          action: golang
          goos: ${{ matrix.goos }}
          goarch: ${{ matrix.goarch }}
          bin_name: "test"
          extra_files: LICENSE  README.md 
~~~

| Variable      | Default             | Example               | Description                                           |
| ------------- | ------------------- | --------------------- | ----------------------------------------------------- |
| `extra_files` | `""`                | `"README.md LICENSE"` | Additional files to include in the compressed package |
| `bin_name`    | `basename $(pwd)`   | `mycli`               | Name of the built binary                              |
| `goos`        | `go env GOHOSTOS`   | `linux`, `windows`    | Target OS for Go build                                |
| `goarch`      | `go env GOHOSTARCH` | `amd64`, `arm64`      | Target architecture                                   |
| `build_flags` | `""`                | `-trimpath`           | Extra build flags for `go build`                      |
| `ldflags`     | `""`                | `-s -w`               | Linker flags for `go build`                           |
| `goamd64`     | `""`                | `v3`                  | `GOAMD64` variant (only for `amd64`)                  |
| `goarm`       | `""`                | `7`                   | ARM version (only for `arm`)                          |
| `gomips`      | `""`                | `softfloat`           | MIPS setting (only for `mips`, `mipsle`, etc.)        |


## GitHub action (golangs upload)

~~~
name: build golangs

on:
  release:
    types: [created]

permissions:
    contents: write
    packages: write

jobs:
  release-linux-amd64:
    name: release linux/amd64
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
    - uses: pkg6/gruploader-action@main
      with:
        github_token: ${{ secrets.GH_TOKEN }}
        action: golangs
        bin_name: test
        extra_files: 
          README.md
          LICENSE
~~~


## GitHub action (node upload)

~~~
name: build node

on:
  release:
    types: [created]

permissions:
    contents: write
    packages: write

jobs:
  release-linux-amd64:
    name: release linux/amd64
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
    - uses: pkg6/gruploader-action@main
      with:
        github_token: ${{ secrets.GH_TOKEN }}
        action: node
        bin_name: web
        extra_files: 
          README.md
          LICENSE
~~~

| Variable            | Default Value     | Example               | Description                                              |
| ------------------- | ----------------- | --------------------- | -------------------------------------------------------- |
| `extra_files`       | `""` (empty)      | `"README.md LICENSE"` | Extra files or directories to copy into the build output |
| `bin_name`          | `"web"`           | `"myapp"`             | Output archive name prefix (e.g., `myapp.tar.gz`)        |
| `npm_install_pkg`   | `"npm install"`   | `"npm ci"`            | Command to install dependencies                          |
| `npm_build_command` | `"npm run build"` | `"vite build"`        | Command to build the project                             |
| `npm_build_dist`    | `"dist"`          | `"build"`             | Directory containing the build output                    |

## GitHub action (node-gh-page)

~~~
name: build node-gh-page

on:
  workflow_dispatch:

permissions:
    contents: write
    packages: write

jobs:
  release-linux-amd64:
    name: release linux/amd64
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
    - uses: pkg6/gruploader-action@main
      with:
        github_token: ${{ secrets.REPO_TOKEN }}
        action: node-gh-page
~~~

| 变量名              | 默认值                      | 说明                                                         |
| ------------------- | --------------------------- | ------------------------------------------------------------ |
| `cname_domain`      | `""`                        | Custom CNAME domain name (CNAME file for GitHub Pages)       |
| `origin_url`        | `""`                        | Git repository address (this value is used first, followed by `git config`, and then GitHub environment variable derivation) |
| `npm_install_pkg`   | `npm ci --legacy-peer-deps` | Install dependency commands                                  |
| `npm_build_command` | `npm run build`             | Build Commands                                               |
| `npm_build_dist`    | `dist`                      | Build product output directory                               |
