FROM golang:alpine AS backend

WORKDIR /src
RUN apk add --no-cache make
COPY go.mod go.sum ./
RUN go mod download
COPY Makefile ./
COPY backend ./backend
RUN make backend-build

FROM node:alpine AS frontend

WORKDIR /src
RUN apk add --no-cache make
COPY package.json yarn.lock ./
RUN yarn install
COPY Makefile tsconfig.json webpack.config.js ./
COPY frontend ./frontend
RUN make frontend-build

FROM alpine

EXPOSE 3455
ENV ADDR="0.0.0.0:3455"
WORKDIR /app
RUN apk add --no-cache make
COPY Makefile ./
COPY --from=backend /src/backend/out/main ./backend/out/
COPY --from=frontend /src/frontend/html ./frontend/html
COPY --from=frontend /src/frontend/css ./frontend/css
COPY --from=frontend /src/frontend/js/out ./frontend/js/out
CMD ["make", "backend-prod"]
