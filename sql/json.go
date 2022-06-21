package sql

import jsoniter "github.com/json-iterator/go"

// json es un paquete que teiene mejor performance que el nativo de Go
var json = jsoniter.ConfigCompatibleWithStandardLibrary
