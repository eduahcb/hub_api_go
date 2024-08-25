FROM golang:1.22.5-alpine AS base

FROM base AS build

WORKDIR /api

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./build/main ./cmd/server/main.go
RUN go build -o ./build/seed ./seeds/seeds.go

FROM base AS application

COPY --from=build api/build ./build
COPY --from=build api/run.sh .
COPY --from=build api/scripts/run_seeds.sh .

RUN ls -a
RUN chmod +x ./run.sh
RUN chmod +x ./run_seeds.sh

# Mode of application
# ENV HUB_API=

# ENV PORT=
# ENV SECRET_KEY=
# ENV DB_URL=
# ENV REDIS_ADDR=

#Expiration time is minutes
# EXPIRATION_TIME=

ENTRYPOINT ["sh", "./run.sh"]