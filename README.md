# Como iniciar el proyecto

Tenemos que abrir dos terminales, una para el backend y otra para el frontend.

## FRONTEND

* En la primera terminal, abriremos la carpeta del frontend con el siguiente comando

`cd frontend-svelte`

* Una vez que la terminal se encuentre posicionada en "frontend-svelte", instalaremos los paquetes necesarios

`npm i`

* Iniciaremos el servidor con:

`npm run dev`

## BACKEND

* En otra terminal, nos tenemos que mover a la carpeta del backend

`cd backend-golang`

* Ya con la terminal posicionada en "backend-golang", vamos a instalar todos los paquetes

`go mod download`

* Despues de instalar todos los paquetes, vamos a compilar el proyecto.

`go build`

* Y por último, iniciar el servidor

`go run .`
