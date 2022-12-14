FROM golang:1.19
WORKDIR /usr/src/app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./src ./src
RUN go build -o ./bin/pomodoro-graph ./src
EXPOSE 3000
CMD [ "./bin/pomodoro-graph" ]
