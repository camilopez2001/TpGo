# TpGo

En nuestro main tenemos nuestra funcion Convertor, la cual se le pasa un string y devuelve una estructura y un error en caso de que lo haya. Para ello creamos una nueva estructura del tipo Result. Luego le asignamos a los atributos de dicha estructura partes definidas del string, comprobando que estas partes cumplan con todas las condiciones:
* Que la variable type sea 2 caracteres letra, ya sea miniscula o mayuscula.
* Que el Length sea 2 numeros int y que esa numero concuerde con la cantidad de caracteres del value
* Que el Value sean carecteres numericos o alfabeticos pero no ambos.
De no cumplir con estas condiciones se notificara el error por pantalla.

En nuestro main_test corremos un json de varios casos en los que puede fallar convertor y mediante un for imprimimos todos los casos. De haber un error este sera notificado.

El test llega a un 89.1% de cobertura.

INTEGRANTES:
Lopez, Camila.
Amici, Cristian.
