FROM scratch

COPY ./artifacts/svc /svc
COPY ./api/service.proto /
COPY ./data/migrations /data/migrations

EXPOSE 8080

CMD ["./svc"]
