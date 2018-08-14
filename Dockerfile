FROM scratch

COPY ./artifacts/svc /svc
COPY ./api/service.proto /

EXPOSE 8080

CMD ["./svc"]
