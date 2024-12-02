FROM golang:1.22-alpine as build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /app/bin/vtech-qms-be

FROM gcr.io/distroless/static-debian12:debug
WORKDIR /app
COPY --from=build /app/bin/. .
EXPOSE 1323
CMD ["./vtech-qms-be"]
