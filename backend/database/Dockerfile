# Use the official MySQL image as the base image
FROM mysql:8.0

# Copy the initialization SQL script to the container.
# The docker-entrypoint-initdb.d directory automatically executed by the official MySQL Docker image when the container is started. This runs init.sql script.
COPY init.sql /docker-entrypoint-initdb.d/

# Expose the MySQL port (3306)
EXPOSE 3306