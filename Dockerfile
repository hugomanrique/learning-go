
# Usa la imagen oficial de MySQL
FROM mysql:latest

# Establece variables de entorno para la configuración de MySQL
ENV MYSQL_DATABASE=mydatabase \
    MYSQL_USER=myuser \
    MYSQL_PASSWORD=mypassword \
    MYSQL_ROOT_PASSWORD=rootpassword

# Copia un archivo SQL (opcional) para inicializar la base de datos
# Si tienes algún archivo .sql o .sh para inicializar, lo puedes copiar a /docker-entrypoint-initdb.d/
# COPY ./init.sql /docker-entrypoint-initdb.d/

# Exponer el puerto MySQL (opcional, por defecto es 3306)
EXPOSE 3306

# Comando por defecto para ejecutar MySQL
CMD ["mysqld"]

# docker build -t mysql-image .
# docker run -d -p 3306:3306 --name my-mysql-container mysql-image