
# Panik

Parse NIK (National Identification Number) Indonesia, this project purpose to extract information from NIK format such as province, city, district, gender and birthdate


## Instalation

Use go get command

```bash
  go get github.com/bicky33/panik
```

Then Import in go project

```bash
  import "github.com/bicky33/panik"
```

## Usage/Examples

```go
parser := panik.NIK{NIK: "Your NIK"}
data, err := parser.Data()

```

## Source Data

Source data such as province code, district, etc. i got from  [here](https://kodewilayah.id/) which based on Permendagri No. 72/2019


