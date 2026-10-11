FROM golang:1.26 AS build
WORKDIR /src
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /hello .

# distroless: sem shell, sem pacote. USER numérico porque o OpenShift
# exige runAsNonRoot e troca o UID por um aleatório de qualquer jeito.
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /hello /hello
USER 65532
EXPOSE 8080
ENTRYPOINT ["/hello"]
