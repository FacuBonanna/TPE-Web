# PROYECTO ESPECIAL PROGRAMACIÓN WEB. 
# Etapa 5: Renderizado del Lado del Servidor con Templ 

En este trabajo reemplazaremos la arquitectura de frontend (HTML/JS/CSS separados) por una aplicación renderizada en el servidor,
utilizando templ para generar las vistas dinámicamente en Go.

## Modelo de la base de datos: 

<img width="1004" height="557" alt="TPE-WEB-DBMODEL-2026-11-08_01-19" src="https://github.com/user-attachments/assets/cc3480d7-0e08-4451-a6e8-a1b43ab8f120" />


## ¿Cómo probar el trabajo? 

1- Clonar el respositorio desde la rama con el siguiente comando 
```
 git clone https://github.com/FacuBonanna/TPE-Web.git
```
1.5- Cambiar la ubicación actual a la ubicación de la carpeta del trabajo (solo en caso de ser necesario) con el siguiente comando 
```
 cd TPE-Web
```
2- Cambiar la rama actual a la rama "tp5" 
```
  git switch tp5
```
3- Para levantar la aplicación, ejecutar el comando:
```
 make inicializacion
```
4- Para visualizar el sistema, ingresar a

http://localhost:8080

nota: si hay algún problema con los comandos tipo 'make', se puede pedir ayuda por consola ingresando el comando 
```
 make --help
```
## Herramientas del proyecto 

-Templ

## Documetos de interés para esta entrega:

-Carpeta views= Ahora que se abandonó la lógica Java script, la caperta views contiene la lógica que hace que el sistema web sea dinámico y funcional. Cada archivo con el formato '*.templ' contiene la lógica de formulario y de lista de manera conjunta

-Components.templ= Dcoumento que contiene métodos comunes a todos los .templ, se decidió esta implemetación para evitar duplicar código y cometer un code smell. 

## Alaraciones y recomendaciones 

-No borrar la carpeta 'views', de lo contrario el sistema no podrá ser usado


