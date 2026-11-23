# PROYECTO ESPECIAL PROGRAMACIÓN WEB. 
# Trabajo Práctico 6 - Interfaces Dinámicas con HTMX

En este trabajo interntaremos evolucionar la aplicación renderizada en el servidor para que se sienta como una aplicación de página única (SPA), utilizando HTMX para eliminar las recargas completas. Además, trataremos pulir la interfaz del sistema Web para que sea más agradable a la vista. 

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
2- Cambiar la rama actual a la rama "tp6" 
```
  git switch tp6
```
3- Para levantar la aplicación (y cerrarla posteriormente), ejecutar el comando:
```
 make inicializacion
```

nota: durante la ejecución de este comando es probable que la terminal se bloquee y pida elegir entre las opciones 'enable watch' y 'detatch'
para poder bajar la base de datos seleccione 'detatch'. Si quiere poder usar la consola mientras el sistema está levantado, ejecute la siguiente línea de comandos: 

```
make levantar
```

4- Para visualizar el sistema, ingresar a

http://localhost:8080

nota: si hay algún problema con los comandos tipo 'make', se puede pedir ayuda por consola ingresando el comando 
```
 make --help
```
## Herramientas del proyecto 

-HTMX
-Templ

## Documetos de interés para esta entrega:

-layout.templ= Contiene toda la lógica HTMX para que la página intercepte acciones del usuario en el HTML, envíe peticiones al servidor y reemplace partes de la página con el HTML nuevo que responde el servidor

-Carpeta views= Es relevante para la entrega porque ahora ca cada entidad se le agrega lógica HTMX con los los atributos hx-post, hx-target y hx-swap. Además, cada entidad es la encargada de su propio botón de delete que se encarga de borrar las instancias sin recargar la página. 

-Main.go= Modificamos los handlers de POST para que después de crear la nueva entidad en la base de datos, vuelva a consultar la lista completa de entidades, la renderice y devuelva únicamente el componente de la lista con los datos actuaizados.



