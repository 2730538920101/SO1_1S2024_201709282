## PASOS PARA COMPILAR EL PROTO

*   Primero debemos instalar el compilador de protobuff para linux con el siguiente comando:

    -   apt install -y protobuf-compiler

*   Inicializar el go mod con el comando:
    -   go mod init "nombre del modulo"

*   Instalar plugins de go para utilizar protoc-gen-go con los siguientes comandos:
    -   go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    -   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    -   export PATH="$PATH:$(go env GOPATH)/bin"

*   Dar permisos de ejecución al script init-proto.sh y ejecutarlo con los comandos:
    -   chmod +x /proto/init-proto.sh
    -   ./init-proto.sh

*   Dentro del script se encuentra el comando base (ejemplo tomado de la documentacion oficial de grpc):
    -   protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto

*   Luego de compilar el proto, es necesario instalar un modulo de go con el comando:
    -   google.golang.org/grpc

*   No olvidar actualizar los modulos de go con el comando:
    -   go mod tidy
