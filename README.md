# ai_dev
Labs from AI DEV 2 course. Each directory is the solution for labs in the selected language.

### test server
directory: test_server

Simple implementation of server that simulates server used during the course by lecturers to verify tasks.

It does not represent the real logic provided by lecturers. It is only a simulator/mock for tests.
```
docker compose run web npm install
docker compose up -d

docker compose stop

docker compose start
```

test server would be available on http://localhost:8080
### Go lang
directory: go_lang
```
go get
go build
chmod +x ai_dev
ai_dev cXXlXX
```
cXXlXX replace with lab number, for example c01l01
