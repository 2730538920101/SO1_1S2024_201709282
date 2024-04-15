#!/bin/bash

# Ruta del archivo .proto
PROTO_FILE="bands.proto"

# Comando para compilar el archivo .proto
protoc \
  --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
  $PROTO_FILE

# Verificar si la compilación fue exitosa
if [ $? -eq 0 ]; then
  echo "Proto compilado exitosamente."
else
  echo "Error al compilar el proto."
  exit 1
fi

