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
COPY Makefile rollup.config.js ./
COPY frontend ./frontend
RUN make frontend-build

FROM alpine

ENV HOST="0.0.0.0"
WORKDIR /app
RUN apk add --no-cache make
COPY Makefile ./
COPY --from=backend /src/backend/out/main ./backend/out/
COPY --from=frontend /src/frontend/static ./frontend/static
COPY --from=frontend /src/frontend/out ./frontend/out
CMD ["make", "backend-prod"]
