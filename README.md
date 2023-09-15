# GOLANG - PRUEBA DE RECONOCIMIENTO DE FRAMES NEGROS EN VIDEOS

La intencion de este projecto es experimentar un sistema de deteccion de videos que contienen muchos o todos los Frames negros, ante lo cual su calidad se consideraria muy baja en casos donde se requiere mostrar contendos media de buena calidad.

Se conecta a un bucket de AWS S3 y analiza una carpeta especifica, al final genera un csv con los resultados obtenidos del analisis.


## LIBRERIAS USADAS
- ffmpeg
- viper
- godotenv
- gorm

Desarrollado por paul@devpwo.com

MIT License

Copyright (c) 2023 Paul Osinga

Se concede permiso, de forma gratuita, a cualquier persona que obtenga una copia de este software y de los archivos de documentación asociados (el "Software"), para tratar el Software sin restricciones, incluyendo, sin limitación, los derechos de uso, copia, modificación, fusión, publicación, distribución, sublicencia y/o venta de copias del Software, y permitir a las personas a quienes se les proporcione el Software que lo hagan, con sujeción a las siguientes condiciones:

El aviso de copyright anterior y este aviso de permiso se incluirán en todas las copias o partes sustanciales del Software.

EL SOFTWARE SE PROPORCIONA "TAL CUAL", SIN GARANTÍA DE NINGÚN TIPO, EXPRESA O IMPLÍCITA, INCLUYENDO, PERO NO LIMITADO A, LAS GARANTÍAS DE COMERCIABILIDAD, APTITUD PARA UN PROPÓSITO PARTICULAR Y NO INFRACCIÓN. EN NINGÚN CASO LOS AUTORES O TITULARES DE LOS DERECHOS DE AUTOR SERÁN RESPONSABLES DE CUALQUIER RECLAMO, DAÑO U OTRA RESPONSABILIDAD, YA SEA EN UNA ACCIÓN DE CONTRATO, AGRAVIO O DE OTRO MODO, DERIVADA DE, FUERA DE O EN CONEXIÓN CON EL SOFTWARE O EL USO O OTRO TIPO DE ACCIONES EN EL SOFTWARE.


