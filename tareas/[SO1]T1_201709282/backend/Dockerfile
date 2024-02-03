# Usa una imagen de Go como base
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos necesarios al contenedor
COPY . .

# Descarga las dependencias del módulo Go (si estás utilizando módulos)
RUN go mod download

# Compila la aplicación Go
RUN go build -o main .

# Expone el puerto en el que la aplicación se ejecutará
EXPOSE 5000

# Comando para ejecutar la aplicación al iniciar el contenedor
CMD ["./main"]
