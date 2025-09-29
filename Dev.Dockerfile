FROM golang:tip-bullseye as builder
WORKDIR /app
COPY . .