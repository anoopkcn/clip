# lazyCLIP

:fish: lazy Command LIne Paper(CLIP) Manager. 

## Usage

### Search

*Simple Usage*
```bash
clip search --string "electron"
```
*Tuning search with more options*
```bash
clip search --source arxive --string "electron scattering" --start 0 --results 5 --prefix "cat" prefix-value "hep-th"
```
One word strings need not to be in double quotes

### Lookup

*Simple Usage*
```bash
clip lookup --doi "DOI/of/a/paper"
```
*Tuning lookup with more options*
```bash
clip lookup --source crossref --doi "my/paper/doi" --bibtex --file mybib.bib
```

