FROM golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM postgres:alpine

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=1
    
COPY init.sql /docker-entrypoint-initdb.d/
COPY --from=build /app/main /bin/app

WORKDIR /app
COPY --chmod=755 entrypoint.sh .

CMD ["/app/entrypoint.sh"]