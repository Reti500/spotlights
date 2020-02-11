# Spotlight project

Prueba de logica, hay que iluminar todas las habitaciones.

Este poryecto esta echo en el lenguaje [Golang], está compilado para los sistemas operativos: 

* OSX
* Linux
* Windows

## Estructura del proyecto:
* main.go   (Código fuente del proyecto).
* in.txt    (Archivo que define la matriz de entrada para hacer el calculo).
* go.mod    (Archivo de dependencias del proyeco).
* go.sum    (Archivo de dependencias del proyeco).

[Golang]: https://golang.org/

## Proyecto

**Problema:** 

Dada una matriz mxn de enteros, calcular la forma más eficiente para iluminar todas casillas donde:

* Un 0 significa que la casilla está libre.
* Un 1 significa que hay una pared.
* Un 2 significa que hay un foco.
* Un 3 significa que la casilla está iluminada.

Ejemplo de matriz de entrada:

    [ 1 0 0 ]
    [ 0 1 0 ]
    [ 0 0 1 ]

Ejemplo de matriz de salida:

    [ 1 0 2 ]
    [ 0 1 0 ]
    [ 2 0 1 ]
    
    
**Solución**

La solución más simple que se me ocurrio fue hacerlo por "fuerza bruta":
    
    [1] Rocorrer la matriz hasta encontrar una casilla con un 0(casilla libre o no iluminada).
        Si no hay casillas, pasamos a el paso [5]
    [2] "Colocar un foco" y calcular cuantas casillas ilumina, se itera a todos lados derecha, izquierda arriba y abajo 
        hasta que encuentre una pared (número 1) o se acaben las habitaciones. Se almacena el resultado en un arreglo 
        llamado "spots".
    [3] Ordenamos el arreglo "spots" y tomamos el que más casillas ilumina.
    [4] Editamos la matriz principal y colocamos un foco(número 2) en la casilla dicha por el spot más grande,
        hacemos el paso [2] pero ahora en vez de calcular marcamos las casillas como iluminadas(numero 3).
    [Repetimos paso 1]
    [5] Iteramos en la nueva matriz y reemplazamos las casillas iluminadas(número 3) por casillas libres(número 0).
    [6] Imprimimos el resultado.
    
## Como correr el proyecto

1. Editar el archivo in.txt y poner la matriz que se quiera evaluar.
2. Usando un compilado:
    * Elejir el archivo spotlights_[os] dependiendo de tu sistema operativo y ejecutarlo; 
    Ejemplo Linux:  ```./spotlights_linux``` 
3. Compilando directo el codigo fuente:
    * Descargar e instalar el lenguaje golang [link](https://golang.org/dl/).
    * En la terminal, moverse al directorio donde está el código fuente y ejecutar el comando: ```go run main.go```

