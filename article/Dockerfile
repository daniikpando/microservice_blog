
# Imagen base
FROM golang:1.9

# Instalación de dependencias
RUN go get github.com/lib/pq
RUN go get github.com/labstack/echo
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/jinzhu/gorm

# Configurar destino del microservicio
COPY . /go/src/github.com/daniel/basic_microservice/article

# Configurar directorio de trabajo
WORKDIR /go/src/github.com/daniel/basic_microservice/article

# Generar el binario ejecutable
RUN go install github.com/daniel/basic_microservice/article

# Punto de entrada
ENTRYPOINT /go/bin/article

# Exponer un puerto
EXPOSE 8080