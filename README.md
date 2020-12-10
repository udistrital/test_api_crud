# test_api_crud

Api de practica para implementar y documentar los lineamienos

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
# parametros de api
TEST_API_CRUD_HTTP_PORT=[Puerto de exposición del API]
TEST_API_CRUD_RUN_MODE=[Modo de ejecución del API]
# paramametros de bd
TEST_API_CRUD_PGUSER=[Usuario de BD]
TEST_API_CRUD_PGPASS=[Contraseña del usaurio de BD]
TEST_API_CRUD_PGHOST=[URL, Dominio o EndPoint de la BD]
TEST_API_CRUD_PGPORT=[Puerto de la BD]
TEST_API_CRUD_PGDB=[Nombre de Base de Datos]
TEST_API_CRUD_PGSCHEMA=[Nombre del Esquema de Base de Datos]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con TEST_API_CRUD_...


### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/test_api_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/test_api_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
TEST_API_CRUD_PGPORT=8080 TEST_API_CRUD_PGHOST=127.0.0.1 TEST_API_CRUD_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile
```shell
# Implementado para despliegue del Sistema de integración continua CI.
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/test_api_crud

#2. Moverse a la carpeta del repositorio
cd test_api_crud

#3. Crear un fichero con el nombre **custom.env**
touch .env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```
### Ejecución Pruebas

Pruebas unitarias
```shell
# Not Data
```
## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| 1 | 2 | 3 |

## Modelo de Datos

[Modelo Relacional test_api_crud]()


## Licencia

This file is part of test_api_crud.

test_api_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

test_api_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with test_api_crud. If not, see https://www.gnu.org/licenses/.
