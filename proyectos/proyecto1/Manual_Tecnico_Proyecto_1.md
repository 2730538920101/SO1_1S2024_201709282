# Manual Técnico Proyecto 1

Creado por: CARLOS JAVIER MARTÍNEZ POLANCO
Fecha de creación: 17 de marzo de 2024 22:02
Etiquetas: Desarrollo

# Modulos de kernel linux

```markdown
### ¿Qué son los módulos de kernel en Linux?

En el núcleo (kernel) de un sistema operativo Linux, los módulos son fragmentos de código que pueden ser cargados y descargados dinámicamente en tiempo de ejecución. Esto permite extender las capacidades del kernel sin necesidad de reiniciar el sistema. Los módulos de kernel pueden añadir soporte para nuevo hardware, sistemas de archivos, protocolos de red, entre otras funcionalidades.

### Funcionamiento de los módulos de kernel:

1. **Compilación**: Los módulos de kernel son compilados por separado del kernel principal. Esto significa que pueden ser desarrollados y distribuidos de forma independiente.

2. **Carga (o inicialización)**: Los módulos se cargan en el kernel utilizando herramientas como `insmod` o `modprobe`. Durante este proceso, el kernel reserva memoria y establece las estructuras de datos necesarias para el módulo.

3. **Utilización**: Una vez cargado, el módulo puede ser utilizado por el kernel según sea necesario. Por ejemplo, un módulo que proporciona soporte para un nuevo dispositivo permitirá que el kernel interactúe con ese dispositivo.

4. **Descarga (o eliminación)**: Los módulos pueden ser descargados del kernel en cualquier momento utilizando `rmmod`. Durante este proceso, se liberan los recursos que el módulo estaba utilizando.

```

### 1. Creación del módulo:

```c
cCopy code
// modulo_ejemplo.c
#include <linux/init.h>#include <linux/module.h>static int __init ejemplo_init(void) {
    printk(KERN_INFO "Hola, este es un ejemplo de módulo de kernel.\n");
    return 0; // Éxito
}

static void __exit ejemplo_exit(void) {
    printk(KERN_INFO "Saliendo del módulo de ejemplo.\n");
}

module_init(ejemplo_init);
module_exit(ejemplo_exit);

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Tu Nombre");
MODULE_DESCRIPTION("Un simple módulo de kernel de ejemplo");
MODULE_VERSION("0.1");

```

### 2. Compilación del módulo:

```bash
bashCopy code
$ sudo apt-get install build-essential linux-headers-$(uname -r)
$ make -C /lib/modules/$(uname -r)/build M=$(pwd) modules

```

### 3. Carga del módulo:

```bash
bashCopy code
$ sudo insmod modulo_ejemplo.ko

```

### 4. Verificación del mensaje de carga:

```bash
bashCopy code
$ dmesg | tail

```

### 5. Eliminación del módulo:

```bash
bashCopy code
$ sudo rmmod modulo_ejemplo

```

### 6. Verificación del mensaje de descarga:

```bash
bashCopy code
$ dmesg | tail

```

Estos pasos ilustran cómo crear, cargar e eliminar un módulo de kernel simple en Linux. Los mensajes de registro pueden ser revisados con el comando **`dmesg`**.

# Docker y Docker Compose

**Docker** es una plataforma de código abierto diseñada para automatizar el despliegue, el desarrollo y la ejecución de aplicaciones en contenedores. Los contenedores permiten empaquetar una aplicación junto con todas sus dependencias en un entorno aislado, garantizando que se ejecute de manera consistente en cualquier entorno.

**Docker Compose** es una herramienta que permite definir y ejecutar aplicaciones Docker de múltiples contenedores. Con Compose, se define un archivo YAML para configurar los servicios de la aplicación. Esto incluye las imágenes de Docker que se utilizarán, así como la configuración de la red, volúmenes, variables de entorno, entre otros.

### **Uso de un Dockerfile para crear una imagen:**

Un **Dockerfile** es un archivo de texto que contiene una serie de instrucciones que Docker utilizará para construir una imagen. Estas instrucciones especifican cómo configurar el entorno dentro del contenedor y qué comandos ejecutar para configurar la aplicación.

### Ejemplo de Dockerfile:

```
DockerfileCopy code
# Seleccionar la imagen base
FROM ubuntu:20.04

# Instalar paquetes necesarios
RUN apt-get update && apt-get install -y \
    nginx \
    && rm -rf /var/lib/apt/lists/*

# Copiar archivos de configuración
COPY nginx.conf /etc/nginx/nginx.conf
COPY index.html /var/www/html/index.html

# Exponer el puerto 80
EXPOSE 80

# Comando por defecto para ejecutar al iniciar el contenedor
CMD ["nginx", "-g", "daemon off;"]

```

### **Ejemplo de uso de Docker Compose:**

Supongamos que queremos utilizar el Dockerfile anterior para crear un contenedor de Nginx y luego ejecutarlo utilizando Docker Compose.

### 1. Crear un archivo **`docker-compose.yml`**:

```yaml
yamlCopy code
version: '3'
services:
  nginx:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:80"

```

### 2. Crear archivos **`nginx.conf`** e **`index.html`**:

**nginx.conf**:

```
nginxCopy code
user www-data;
worker_processes auto;
pid /run/nginx.pid;
events {
    worker_connections 768;
}
http {
    sendfile on;
    tcp_nopush on;
    tcp_nodelay on;
    keepalive_timeout 65;
    types_hash_max_size 2048;
    include /etc/nginx/mime.types;
    default_type application/octet-stream;
    server {
        listen 80;
        server_name localhost;
        location / {
            root /var/www/html;
            index index.html;
        }
    }
}

```

**index.html**:

```html
htmlCopy code
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Hello Docker</title>
</head>
<body>
    <h1>Hello Docker!</h1>
    <p>This is a Dockerized Nginx web server.</p>
</body>
</html>

```

### 3. Construir y ejecutar el contenedor con Docker Compose:

```bash
bashCopy code
$ docker-compose up -d

```

### 4. Acceder a la aplicación:

Abre un navegador web y visita **`http://localhost:8080`** para ver la página web servida por Nginx en el contenedor Docker.

# Como correr el proyecto?

## Paso 1

Copiar el archivo env-sample.txt  en el directorio raiz del proyecto para poder setear las variables de entorno utilizadas en el docker compose. Se puede usar el comando cp o los comandos cat y echo con salida stdin a un nuevo archivo .env

cat env_sample.txt > .env

No olvidar que la variable NODE_ENV debe ser seteada como “production” para que al hacer el build tome el valor del docker compose y no del arcivo .env, además para eso tambien se usaron los archivos especiales .dockerignore.

## Paso 2

Darle permisos de ejecusión a los scripts en la carpeta raiz con el comando chmod, estos scripts al ser ejecutados tienen la capacidad de inicializar los modulos y de eliminarlos dependiendo de cual se ejecute.

sudo chmod +x init_modules.sh

sudo chmod +x delete_modules.sh

## Paso 3

Ejecutar el script con el comando:

./init_modules.sh

## Paso 4

Ejecutar el comando de inicialización de los contenedores de docker, utilizando las imagenes previamente cargadas con su respectiva versión en el repositorio de docker hub. Utilizando el comando:

docker compose up -d 

## Paso 5

Abrir el navegador con la dirección: http://localhost

Si corremos el programa desde una maquina virtual sea en la nube o sea en virtualbox o vmware, utilizamos la ip publica que nos proveen los servicios para acceder.

## Paso 6

Para eliminar los modulos ejecutar el script en la carpeta raiz con el comando:

./delete_modules.sh