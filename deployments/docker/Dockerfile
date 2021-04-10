FROM golang:latest AS dependencies
ENV NAME "app"
WORKDIR /opt/${NAME}
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM dependencies AS build
ENV NAME "app"
WORKDIR /opt/${NAME}
COPY . .
RUN ["make", "build"]

FROM alpine
ARG ENV
ARG VERSION
ENV NAME "app"
WORKDIR /opt/${NAME}
COPY --from=build /opt/${NAME}/bin/${NAME} ./${NAME}
CMD ./${NAME} ${CONFIG_PATH}