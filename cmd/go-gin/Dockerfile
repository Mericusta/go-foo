FROM public.ecr.aws/docker/library/golang
WORKDIR /app
COPY . .
RUN go build -o gin-app cmd/go-gin/main.go
CMD ["./gin-app"]
EXPOSE 8080