FROM scratch
# Copiar los certificates TLS de la imagen alpine, que permitirán que la aplicación pueda hacer peticiones HTTP.
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/pkmcli"]
COPY pkmcli /
