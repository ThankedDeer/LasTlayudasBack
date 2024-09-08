# Establecer la imagen base como la imagen oficial de Go 1.22.0
FROM golang:1.22.4 AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos necesarios para descargar los módulos Go
COPY ./go.mod ./go.sum ./

# Descargar los módulos Go
RUN go mod download

# Copiar el resto del código fuente de la aplicación desde dos niveles más abajo
COPY . .

# Cambiar al directorio /app/cmd
WORKDIR /app/cmd

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping .

# Etapa de construcción final
FROM alpine:latest

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el binario de la aplicación compilada desde la etapa de construcción anterior
COPY --from=builder /docker-gs-ping .

# Copiar el archivo .env al contenedor
COPY app.env .

# Exponer el puerto en el que la aplicación escucha
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./docker-gs-ping"]