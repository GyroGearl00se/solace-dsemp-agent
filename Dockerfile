FROM golang:1.24 AS build

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=1 GOOS=linux go build -o /build/solace-dsemp-agent

FROM gcr.io/distroless/cc AS final

COPY --from=build /build/solace-dsemp-agent /solace-dsemp-agent

ENTRYPOINT ["/solace-dsemp-agent"]