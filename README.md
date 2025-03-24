# Banking System Microservices

## ¿Qué es Banking System API?
Es un sistema bancario basado en microservicios que permite la gestión de cuentas bancarias y transferencias entre ellas. La aplicación está desarrollada en Golang utilizando una arquitectura hexagonal (Ports and Adapters) y consta de dos servicios independientes:

Servicio de Cuentas (account-service): Maneja la creación y consulta de cuentas bancarias.

Servicio de Transacciones (transaction-service): Procesa transferencias entre cuentas y registra el historial de transacciones.

Los servicios se comunican mediante API REST y están diseñados para usar PostgreSQL como base de datos y Kafka para mensajería asíncrona (pendiente de implementación). La autenticación con JWT está en desarrollo. La aplicación está orquestada con Docker Compose para facilitar su ejecución y despliegue.

---
## Requerimientos

- Go 1.21 o superior
- Docker y Docker Compose
- PostgreSQL (incluido en Docker Compose)

---
## Instalación

#### 1. Clonar proyecto: 

```bash
https://github.com/castellarmartinez/bank-system.git
cd bank-system
```

#### 2. Instalar dependencias de Go

```
go mod download
```

#### 3. Configurar variables de entorno

Crea un archivo .env en la raíz del proyecto basado en .env.example y ajusta los valores según tu entorno:

```
DB_USER=user
DB_PASSWORD=password
DB_PORT=5432
DB_NAME=banksystem
DB_SSLMODE=disable

HOST=localhost
ACCOUNT_SERVICE_PORT=8081
TRANSACTION_SERVICE_PORT=8082
BFF_SERVICE_PORT=8083

API_KEY=vectora_api_key
JWT_SECRET=jwt_secret
JWT_EXPIRATION=3600000
```

#### Configurar Docker Compose

El archivo docker-compose.yml incluye los servicios account-service, transaction-service, y postgres. Asegúrate de que el archivo db/init.sql esté configurado con las tablas necesarias.

---
## ¿Cómo ejecutar la aplicación?

### Usando Docker Compose

1. Levantar los servicios:

```
docker-compose up --build
```

Esto construirá las imágenes de los servicios y ejecutará los contenedores. El account-service estará disponible en http://localhost:...., el transaction-service en http://localhost:.... y el bff http://localhost:....

2. Detener los servicios:

```
docker-compose down
```

### Manualmente (sin Docker)

1. Inicia PostgreSQL localmente y ejecuta db/init.sql para crear las tablas.

2. Desde account-service/cmd:

```
go run main.go
```

2. Desde transaction-service/cmd:

```
go run main.go
```

---
## ¿Cómo usar la API?

### Servicio de Cuentas (account-service)

#### 1. Crear una cuenta
 
Ruta: POST /accounts
Request: 

```
{
    "nombre": "Juan Pérez",
    "saldo_inicial": 1000
}
```


Response: 

```
{
    "id": 1,
    "nombre": "Juan Pérez",
    "saldo": 1000
}
```


Ejemplo:

```
curl -X POST -H "Content-Type: application/json" -d '{"nombre":"Juan Pérez","saldo_inicial":1000}' http://localhost:8081/accounts
```


#### 2. Consultar saldo de una cuenta

Ruta: GET /accounts/{id}

Response: 

```
{
    "saldo": 1000
}
```


Ejemplo:

```
curl -X GET http://localhost:8081/accounts/1
```

### Servicio de Transacciones (transaction-service)

#### 1. Transferir dinero
 
Ruta: POST /transactions

Request: 

```
{
    "from_account": 1,
    "to_account": 2,
    "monto": 500
}
```


Response: 

```
{
    "status": "success",
    "transaction_id": 1001
}
```


Ejemplo:

```
curl -X POST -H "Content-Type: application/json" -d '{"from_account":1,"to_account":2,"monto":500}' http://localhost:8082/transactions
```


#### 2. Consultar historial de transacciones

Ruta: GET /transactions/{account_id}

Response: Lista de transacciones asociadas a la cuenta:


```
[
    {
        "id": 1001,
        "from_account": 1,
        "to_account": 2,
        "monto": 500,
        "timestamp": "2025-03-24T10:00:00Z",
        "status": "pending"
    }
]
```


Ejemplo:

```
curl -X GET http://localhost:8082/transactions/1
```

## Construido con: 

- [Go](https://go.dev/) - Lenguaje de programación eficiente y concurrente.
- [PostgreSQL](https://www.postgresql.org/) - Sistema de gestión de bases de datos relacional.
- [Docker](https://www.docker.com/) - Plataforma de contenedores para orquestación.
- [Docker Compose](https://docs.docker.com/compose/) -  Herramienta para definir y ejecutar aplicaciones multi-contenedor.
- [JWT](https://jwt.io/) - Autenticación basada en tokens.

---
## Autor 
**David Castellar Martínez** [[GitHub](https://github.com/castellarmartinez/)]
[[LinkedIn](https://www.linkedin.com/in/castellarmartinez/)]

