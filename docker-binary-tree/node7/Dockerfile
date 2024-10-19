# Node 7 - PostgreSQL Database for Node 3 (Python)
FROM postgres:13

# Set environment variables
ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=dbnode7

# Copy the init script if you have any SQL to initialize
COPY init-db.sql /docker-entrypoint-initdb.d/

EXPOSE 5432