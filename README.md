# lazyCLIP
[![GPL 3](https://img.shields.io/badge/license-GPLv3-blue.svg)](COPYING)

:fish: lazy Command LIne Paper(CLIP) Manager. 

**Only few of the possible options and commands are listed here, Please refer to wiki for a full list of features**

## Usage

```bash
clip [-version][-help] <commands> [<options>]
```
**Note:** Multiple word command-options should be wrapped in double quotes. Commands such as `search`, `lookup` and `add` (for `add` and `search` when `-source` flag is given) require an active internet connection. A flag can be given with a `-` or `--` (ex: `-string` is the same as `--string`). Options to a flag can be separated by a `space` or an `=` (ex: `-string "electron"` is the same as `-string="electron"`)

### Search

*Simple Usage*
```bash
clip search -string "electron"
```
*Tuning search with more options*
```bash
clip search -source arxive -string "electron scattering" -start 0 \
-results 5 -prefix "cat" -prefix-value "hep-th"
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

### Add

### Rm

### Mv

### Commit

### Log

### Batch
