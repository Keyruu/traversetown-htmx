FROM node:18.20.3-slim as tailwind
WORKDIR /app
COPY . .
RUN npm install
RUN npx tailwindcss -i ./main.css -o ./pb_public/main.css

FROM golang:1.22.4-alpine as go-builder
WORKDIR /app
COPY . .
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go build -o out

FROM alpine:latest
WORKDIR /app

RUN apk add curl wget

COPY --from=go-builder /app/out /app/out
COPY --from=tailwind /app/pb_public /app/pb_public

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 CMD [ "curl localhost:8090" ]

EXPOSE 8090
ENTRYPOINT [ "/app/out" ]
CMD ["serve",  "--http=0.0.0.0:8090"]