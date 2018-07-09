

# Creación basica de un microservicio

Microservicio para la gestión de articulos


## Requerimientos

1. Docker 18.03
2. Docker compose 1.21

## Iniciar servicios

Posicionarse en la carpeta principal donde se encuentra el archivo `docker-compose.yml` y ejecutar el siguiente comando:

```
docker-compose up -d
```

Para detener los servicios: 
```
docker-compose stop
```
Para reconstruir los servicios
```
docker-compose up --build
```

## Documentación

Este microservicio se encarga de manejar la gestión de articulos en la aplicación.

### Endpoints 

1. Mostrar todos los articulos:
```
GET /api/v1/articles
```
2. Mostrar un articulo por su id:
```
GET /api/v1/articles/:id
```

3. Crear un articulo: 
```
POST /api/v1/articles
```

4. Actualizar un articulo: 
```
PUT /api/v1/articles/:id
```

5. Eliminar un articulo: 
```
DELETE /api/v1/articles/:id
```

