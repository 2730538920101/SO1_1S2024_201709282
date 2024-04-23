# Dockerfile
FROM mongo:4.4

# Copiar el script de inicialización al directorio /docker-entrypoint-initdb.d/
COPY ./init-db.js /docker-entrypoint-initdb.d/

# Establecer permisos para el script
RUN chmod +x /docker-entrypoint-initdb.d/init-db.js

EXPOSE 27017