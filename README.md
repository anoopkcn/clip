# lazyCLIP
[![GPL 3](https://img.shields.io/badge/license-GPLv3-blue.svg)](COPYING)

lazy Command LIne Paper(CLIP) Manager. 

**Only few of the possible options and commands are listed here, Please refer to wiki for a full list of features**

## Installation

If you haven't setup Go before :astonished:, you need to first set a `GOPATH` (see [https://golang.org/doc/code.html#GOPATH](https://golang.org/doc/code.html#GOPATH)).

To fetch and build the code:

    $ go get github.com/strivetobelazy/lazyclip

This will also build the command line tool `clip` into `$GOPATH/bin` (assumed to be in your `PATH` already).

## Usage

```bash
clip [-version][-help] <commands> [<options>]
```
**Note:** Multiple word command-options should be wrapped in double quotes. Commands such as `search`, `lookup` and `add` (for `add` and `search` when `-source` flag is given: please refer to [supported sources](Supported sources)) require an active internet connection. A flag can be given with a `-` or `--` (ex: `-string` is the same as `--string`). Options to a flag can be separated by a `space` or an `=` (ex: `-string "electron"` is the same as `-string="electron"`)

## Commands

### Search

*Simple Usage*
```bash
clip search -string "electron"
```
*Tuning search with more options*
```bash
clip search -source arxive -string "electron scattering" -offset 0 \
-results 5 -filter "cat" -prefix "hep-th"
```
One word strings need not to be in double quotes
**Note:** For `search`, `-source` flag is active on default with value `"arxiv"`

### Lookup

*Simple Usage*
```bash
clip lookup -doi "DOI/of/a/paper"
```
*Tuning lookup with more options*
```bash
clip lookup -source crossref --doi "my/paper/doi" --bibtex --file mybib.bib
```

### pdf
lazyclip by default doesn't have any external dependencies. However if you compile lazyclip with pdf support it can perform additional tasks and operations on local pdf files(see wiki). For this functionality lazyclip relies on `poppler-utils` which is an open source package and can be installed using the following:
for Linux:
```bash
sudo apt-get install poppler-utils
```
for Mac:
```bash
brew install poppler-utils
```
### Add

### Rm

### Mv

### Commit

### Log

### Batch

### Supported sources
The `-source` flag takes following values:

- `crossref` : [CrossRef](https://www.crossref.org/), an exhaustive academic search engine
- `arxiv` : [arXiv](https://arxiv.org/), an archive of pre-prints in various scientific fields
- `dblp` : [DBLP](https://dblp.uni-trier.de/), a database of Computer Science publications
- `doi.org` : [doi.org](http://www.doi.org/), a DOI resolver
- `dissemin` : [Dissemin](https://dissem.in/), a database tracking the open access status of scholarly articles
