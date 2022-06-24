# Bücherliste  📚
![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![Docker](https://badgen.net/badge/icon/docker?icon=docker&label)](https://https://docker.com/)
[![Build Status](https://travis-ci.com/alichtman/shallow-backup.svg?branch=master)](https://travis-ci.com/alichtman/shallow-backup)


## ¿Qué es? 🤔

Proyecto para generar un API instantáneo de almacenamiento de libros en listas de deseos. Las búsquedas de libros se hacen usando el [API de Google](https://developers.google.com/books).

### ¿Qué es significa "Bücherliste"? 🤨

Significa "lista de libros", es Alemán 🤓 🇩🇪



## ¿Instantáneo?, WOW 🤯

Claro, para ejecutarlo debes tener instalado como mínimo:

1. [Docker](https://docs.docker.com/desktop/#download-and-install) ^20.10.16
2. [Docker Compose](https://docs.docker.com/compose/install/) ^v2.6.0
3. [Lolcat](https://github.com/busyloop/lolcat/) ^100.0.1 (Solo por diversión)

Lolcat es para sistemas Unix/Linux, pero en Windows creo que puedes intalar shells de Linux y ejecutar comandos allí si no pues ni modo 🤷🏽‍♂️.

Para el jugar con el API puedes usar [Postman](https://www.postman.com/downloads/), en él importas la [descripción del API](https://github.com/kalmecak/bucherliste/blob/main/apidescription.yml) y te generará las coleciones necesarias, también puedes descargarte la [colección de peticiones](https://www.getpostman.com/collections/612f64002e25dc9b64a9).
**No olvides:** crear tu colección de variables de entorno en postman, sólo necesitarás _host_ (localhost:8080) y _token_ (JWT token).

También puedes abrir el [documento](https://github.com/kalmecak/bucherliste/blob/main/apidescription.yml) en [Swagger Editor](https://editor.swagger.io/) y se te generará la descripción de los servicios.

## Instalación 📦

Sigue la receta mágica:

**Clona el repositorio en tu directorio de trabajo 📋**
```bash
git clone git@github.com:kalmecak/bucherliste.git
```
**Entra en el directorio**
```bash
cd bucherliste
```
**Escoje la pastilla roja 💊✨**
```bash
docker-compose build | lolcat &&  docker-compose up | lolcat
```

NOTA: Si no tienes lolcat, puedes usar el comando:

```bash
docker-compose build && docker-compose up
```

Después de haber ejecutado el último comando, verás las magia surgir en tu terminal:

Puedes empezar a generar peticiones después de ver este mensaje:
```bash
golang-docker-web  |
golang-docker-web  |  ┌───────────────────────────────────────────────────┐
golang-docker-web  |  │                   Bücherliste                     │
golang-docker-web  |  │                   Fiber v2.34.1                   │
golang-docker-web  |  │               http://127.0.0.1:8080               │
golang-docker-web  |  │       (bound on host 0.0.0.0 and port 8080)       │
golang-docker-web  |  │                                                   │
golang-docker-web  |  │ Handlers ............ 35  Processes ........... 1 │
golang-docker-web  |  │ Prefork ....... Disabled  PID ................. 1 │
golang-docker-web  |  └───────────────────────────────────────────────────┘
golang-docker-web  |

````

**Ejoy!**

## Jugar con el API 🎮

Ya que tengas los contenedores listos, puedes empezar a generar peticiones, te recomiendo usar primero los paths
para crear tu usuario:

* [POST]  http://localhost:8080/signup

y luego crear tu sesión:

* [POST]  http://localhost:8080/login

Después ya podrás jugar con los otros paths:

* Buscar libros: [GET]  http://localhost:8080/books?a=camilla&t=hielo&p=maeva&key=AIzaSyAqlSYDVik9vOBuLLhpIK_TNv7bh-VbHrk
* Crear wishlist: [POST]  http://localhost:8080/wishlist
* Lista de wishlists [GET]  http://localhost:8080/wishlists
* Contenido de una wishlist: [GET]  http://localhost:8080/wishlist/:id
* Agregar ó quitar libros: [PUT]  http://localhost:8080/wishlist/:id
* Eliminar wishlists: [DELETE]  http://localhost:8080/wishlist/:id


## Herramientas utilizadas 🔧

### Base de datos 💾

**MySQL**

Opté por esta BD porque es de uso frecuente, fácil de implementar, ya tengo experiencia con ella y maneja bien los UUIDs.

**UUID**

En mi experiencia, el uso de estos tipos de identificadores tiene beneficios que permiten estar migrando al data wharehouse los registros más antiguos sin problemas
de replicación de IDs cuando se hacen los _truncates_ de las tablas. Además de que al trabajar con estos IDs en APIs es mejor ya que no dan indicios de los números de registros que tiene la BD.

Sí, es cierto que tienen menor rendimiento que los INTs, pero ese rendimiento se puede mantener en óptimo teniendo políticas de limpieza de las tablas cada cierto tiempo para mover los registros antiguos a otra tabla o base de datos.

### Golang 💻
Versión 1.18

### [Fiber 🔌](https://docs.gofiber.io/)

Me gusta este Framework porque es fácil de usar, está basado en ExpressJS en el cual
ya tenía experiencia y se encuentra entre los más rápidos.

### [GORM 💾](https://docs.gofiber.io/)

Es un ORM para Golang que he usado en muchos proyectos por la facilidad de uso, la implemntación rápida y contiene muchas funcionalidades que permiten customizar el funcionamiento, así como la rápida creación de tablas basadas en las estructuras del
proyecto.


## Estructura del proyeto

Usé una aproximación a la estructura de proyecto recomendada por las buenas prácticas de Golang.

```bash
.
├── api
│   ├── books
│   ├── user
│   ├── validate
│   └── wishlist
├── cmd
│   └── migration
├── common
├── config
├── docker
├── environment
├── google
└── sql
```

**root**

En la raís del proyecto viven los archivos de configuración generales de Go, readme, docker-compose, configuración del editor y la descripción del API entre otros.

**/api**

Dentro de _api_ se encuentran los paquetes que va a estar funcionando como los manejadores de las peticiones, así como los middlewares que ejecutan validaciones en los requests.

**cmd**

Normalmente aquí pongo los scripts que se ejecutan para hacer "cosas" del proyecto que no nesesariamente son peticiones, como por ejemplo, la migración de la base de datos.

**common**

Es el paquete que contiene lógica y estructuras que se usan en diversas partes del proyecto.

**config**

Contiene las configuraciones perse del servidor, content-types, validaciones de sesión, CORS policies, etc.

**docker**

Tiene los Dockerfiles de los contenedores que se usan en el proyecto.

**environment**

Es mi paquete para verificar que todas las variables de entorno necesarias en el proyecto se encuentran definidas.

**google**

Paquete que tiene la lógica de conexión y solicitud con Google Books API.

**sql**
Paquete que tiene las estructuras de la base de datos y lógica relacionada al manejo de los datos.
