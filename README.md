# CLIP : "Under Development: Users-please come back, developers: fork!!"
[![GPL 3](https://img.shields.io/badge/license-GPLv3-blue.svg)](LICENSE)

Command LIne Papers(CLIP), a reference manager.

**Only few of the possible options and commands are listed here, Please refer to wiki for a full list of features**

## Installation

If you haven't setup Go before :astonished:, you need to first set a `GOPATH` (see [https://golang.org/doc/code.html#GOPATH](https://golang.org/doc/code.html#GOPATH)).

To fetch and build the code:

    $ go get github.com/strivetobelazy/clip/...

This will also build the command line tool `clip` into `$GOPATH/bin` (assumed to be in your `PATH` already).

## Usage

```bash
clip [-version][-help] <commands> [<options>]
```
**Note:** Multiple word command-options should be wrapped in double quotes. Commands such as `search` and `add` (for `add` when `-source` flag is given: please refer to [supported sources](Supported sources)) require an active internet connection. A flag can be given with a `-` or `--` (ex: `-string` is the same as `--string`). Options to a flag can be separated by a `space` or an `=` (ex: `-string "electron"` is the same as `-string="electron"`)

## Commands

### search

*Simple Usage*
```bash
clip search -string "electron"
```
*Tuning search with more options*
```bash
clip search -source arxive -string "electron scattering" -offset 0 \
-results 5 -filter "cat" -prefix "hep-th"
```
*findout more details about a specific information(ex:doi, title, author, etc,.)*
```bash
clip search -source crossref -string "doi/of/any/paper" -match doi
```
*get information in a bibtex format and append it to a file*
```bash
clip search -string "title-of-a-paper" -match title -bibtex -file mybib.bib
```
One word strings need not to be in double quotes

**Note:** For `search`, `-source` flag is active on default with value `"arxiv"`

### lookup

`lookup` command performs all the operations of `search` command but on a local `clip` repository or a folder.

*Simple Usage*
```bash
clip lookup -string "DOI/of/a/paper" -match doi
```
*Tuning lookup with more options*
```bash
clip lookup -string "title-of-a-paper" -bibtex -file mybib.bib
```
### init

### add

### rm

### mv

### log

### batch

## Optional functionality

clip by default doesn't have any external dependencies. However the following enhancements are provided upon compiling clip with additional external packages

### pdf

Compiling clip with pdf support will add `pdf` command which perform additional tasks and operations on local pdf files(see wiki). For this functionality lazyclip relies on `poppler-utils` which is an open source package and can be installed using the following:

for Linux:
```bash
sudo apt-get install poppler-utils
```
for Mac:
```bash
brew install poppler-utils
```

### fzf

Compiling clip with fzf support will bring in compatibility with fzf package

## Additional information

### Supported sources
The `-source` flag takes following values:

- `google` : [google](www.google.com), if you couldn't find anything in any other sources.
- `crossref` : [CrossRef](https://www.crossref.org/), an exhaustive academic search engine
- `arxiv` : [arXiv](https://arxiv.org/), an archive of pre-prints in various scientific fields
- `dblp` : [DBLP](https://dblp.uni-trier.de/), a database of Computer Science publications
- `doi.org` : [doi.org](http://www.doi.org/), a DOI resolver
- `dissemin` : [Dissemin](https://dissem.in/), a database tracking the open access status of scholarly articles

Note: If you like to see support for an other source, please open an [isuue](https://github.com/strivetobelazy/clip/issues) or contact us via email.
