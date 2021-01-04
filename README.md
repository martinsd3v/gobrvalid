# gobrvalid

Golang package for validate brazilian documents and dates

#### Installation
Make sure that Go is installed on your computer.
Type the following command in your terminal:
```bash
go get github.com/martinsd3v/gobrvalid
```

#### Import package in your project
Add following line in your `*.go` file:
```go
import "github.com/martinsd3v/gobrvalid"
```
If you are unhappy to use long `govalidator`, you can do something like this:
```go
import (
  valid "github.com/martinsd3v/gobrvalid"
)
```

#### List of functions:
```go
func IsCPF(cpf string) (isValid bool, clean string)
func IsCNPJ(cnpj string) (isValid bool, clean string)
func IsCPForCNPJ(doc string) (isValid bool, clean string)
func IsCNH(cnh string) (isValid bool, clean string)
func IsVehiclePlate(plate string) bool
func IsVehicleRenavam(renavam string) (isValid bool, clean string)
func IsDate(date string) (isValid bool, dateTime time.Time)
```

###### PT-BR

Pacote em golang para validação de documentos e datas brasileiras

#### Instalação
Certifique-se de que Go está instalado em seu computador.
Digite o seguinte comando em seu terminal:

```bash
go get github.com/martinsd3v/gobrvalid
```

#### Importar pacote em seu projeto
Adicione a seguinte linha em seu arquivo `*.go`:
```go
import "github.com/martinsd3v/gobrvalid"
```
Se você não gosta de usar o `gobrvalid` longo, pode fazer algo assim:
```go
import (
  valid "github.com/martinsd3v/gobrvalid"
)
```

#### Lista de funções:
```go
func IsCPF(cpf string) (isValid bool, clean string)
func IsCNPJ(cnpj string) (isValid bool, clean string)
func IsCPForCNPJ(doc string) (isValid bool, clean string)
func IsCNH(cnh string) (isValid bool, clean string)
func IsVehiclePlate(plate string) bool
func IsVehicleRenavam(renavam string) (isValid bool, clean string)
func IsDate(date string) (isValid bool, dateTime time.Time)
```