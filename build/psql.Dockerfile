FROM postgres:11
COPY build/init.sql /docker-entrypoint-initdb.d/

# not for prod use
ENV POSTGRES_HOST_AUTH_METHOD=trust