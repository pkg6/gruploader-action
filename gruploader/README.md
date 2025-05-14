## Golang Install

~~~
go install github.com/pkg6/gruploader-action/gruploader@latest
~~~

## Usage

- help

```
Usage of ./gruploader:
  -cloud string
        Which cloud do you use for git? (default "github")
  -f string
        File path to upload.
  -overwrite
        Overwrite release asset if it's already exist.
  -repo string
        repo, e.g., 'pkg6/gruploader-action'.
  -retry uint
        How many times to retry if error occur. (default 2)
  -tag string
        Git tag to identify a  Release in repo.

```
## example

~~~
export GITHUB_TOKEN="ghp_**************"
gruploader -f gruploader.tar.gz --repo=pkg6/gruploader --tag=v0.0.1 --overwrite=true
~~~
