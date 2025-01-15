#!/bin/sh

# Gerar documentação
gobiru -output /app/docs/routes.json \
       -openapi /app/docs/openapi.json \
       -title "API Documentation" \
       -description "API documentation generated by Gobiru" \
       -framework mux \
       /app/routes.go

# Iniciar o servidor
exec /usr/local/bin/server 