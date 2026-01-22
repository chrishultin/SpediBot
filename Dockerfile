FROM node:25 AS frontend_build

WORKDIR /app

COPY frontend-old/package.json package.json
COPY frontend-old/package-lock.json package-lock.json

RUN npm install

COPY frontend-old/ .

RUN npm run build

FROM golang:1.25-alpine

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download

COPY ./backend/ ./backend/
COPY ./migrations/ ./migrations/
COPY ./ui-embed/ ./ui-embed/

COPY --from=frontend_build /app/dist/spa/ ./ui-embed/

RUN go build -v -o /usr/local/bin/app ./backend/exec/main.go

EXPOSE 80

CMD ["app", "serve", "--http=0.0.0.0:80"]