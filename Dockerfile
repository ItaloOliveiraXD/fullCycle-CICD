FROM amd64/golang

WORKDIR /app

COPY . .

RUN go build -o math

CMD [ "./math" ]