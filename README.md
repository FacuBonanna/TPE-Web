# PROYECTO ESPECIAL PROGRAMACIÓN WEB. Etapa 3: Exponiendo la lógica como API. 

EL objetivo de esta entrega es construir una API REST para gestionar las entidades de nuestra aplicación para poder conectar la lógica de negocio con la
capa de datos.

## Modelo de la base de datos: 

<img width="1004" height="557" alt="TPE-WEB-DBMODEL-2026-09-24_10-52" src="https://github.com/user-attachments/assets/22cddfb9-52ff-4e93-b1c1-8dec8270a10e" />

## ¿Cómo probar el trabajo? 

1- Clonar el respositorio desde la rama con el siguiente comando 
```
 git clone https://github.com/FacuBonanna/TPE-Web.git
```
1.5- Cambiar la ubicación actual a la ubicación de la carpeta del trabajo (solo en caso de ser necesario) con el siguiente comando 
```
 cd TPE-Web
```
2- Cambiar la rama actual a la rama "tp3" 
```
  git switch tp3
```
3- Para la API REST, ejecutar el comando:
```
  make test_api
```
3.5- Si se quiere probar la base de datos por separado, ejecutar el comando: 
```
  make test_base
```

nota: si hay algún problema con los comandos tipo 'make', se puede pedir ayuda por consola ingresando el comando 
```
 make --help
```
## Herramientas del proyecto 

 - Server Docker
 - Make
 - PostgreSQL
 - Sqlc

## Documetos de interés para esta entrega:

A pesar de que todos las partes del código son importantes para que este proyecto pueda funcionar, estos son los elementos que tuvimos que modificar para poder concretar la entrega: 

- resquests.hurl: ///
- Makefile:
- Carpeta src: 


## Aclaraciones

-Los tests ya están integrados en la lógica del proyecto, se recomienda no remover el archivo requests.hurl del directorio que contiene el archivo. ya que de otra forma el make test no funcionará



