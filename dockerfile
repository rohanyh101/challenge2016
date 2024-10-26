FROM debian:stable-slim

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir /app

COPY ./bin/qube /app

COPY .env /app/.env

COPY ./data/cities.csv /app/data/cities.csv

WORKDIR /app

CMD ["/app/qube"]