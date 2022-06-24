# Configuración la imagen de la base de datos
FROM mysql:oracle

# Copiamos la configuración del collate para la base de datos
COPY ./docker/custom.cnf /etc/mysql/conf.d/custom.cnf
