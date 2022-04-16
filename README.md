# Usage

Project contains Taskfile with available actions. To use Taskfile, you need to install task:

```bash
brew install go-task/tap/go-task
```

More information: https://taskfile.dev/#/installation

To view all available tasks just type:

```bash
task
```

# Development environment

To run development environment run:

```bash
task dev
```

# Tests

To execute tests, run:

```bash
task test
```

# Production
To run project for production, use:

```bash
docker run -p 8080:8080  drymek/spacebook:latest
```

# Usage examples

## Book
```bash
echo '{ "id": "1234", "firstname": "John", "lastname": "Doe", "gender": "Male", "birthday": "2000-07-21", "launchpadID": "5e9e4501f5090910d4566f83", "destinationID": "Mars", "launchDate": "2022-01-17"}' | http POST :8080/bookings
HTTP/1.1 200 OK
Access-Control-Allow-Headers: Origin, Content-Type
Access-Control-Allow-Methods: GET, POST, OPTIONS
Access-Control-Allow-Origin: *
Content-Length: 14
Content-Type: text/plain; charset=utf-8
Date: Sat, 16 Apr 2022 01:44:43 GMT

{
    "id": "1234"
}
```

## Booking error
```bash
$ echo '{ "id": "1234", "firstname": "John", "lastname": "Doe", "gender": "Male", "birthday": "2000-07-21", "launchpadID": "5e9e4502f509094188566f88", "destinationID": "Asteroid Belt", "launchDate": "2022-07-01"}' | http POST :8080/bookings

HTTP/1.1 400 Bad Request
Access-Control-Allow-Headers: Origin, Content-Type
Access-Control-Allow-Methods: GET, POST, OPTIONS
Access-Control-Allow-Origin: *
Content-Length: 75
Content-Type: text/plain; charset=utf-8
Date: Sat, 16 Apr 2022 13:39:46 GMT

{
    "error": "booking service error: SpaceX already has a launch on this date"
}
```

## Get all bookings
```bash
http :8080/bookings
HTTP/1.1 200 OK
Access-Control-Allow-Headers: Origin, Content-Type
Access-Control-Allow-Methods: GET, POST, OPTIONS
Access-Control-Allow-Origin: *
Content-Length: 192
Content-Type: text/plain; charset=utf-8
Date: Sat, 16 Apr 2022 01:46:35 GMT

{
    "items": [
        {
            "birthday": "2000-07-21",
            "destinationID": "Mars",
            "firstname": "John",
            "gender": "Male",
            "id": "1234",
            "lastname": "Doe",
            "launchDate": "2022-01-17",
            "launchpadID": "5e9e4501f5090910d4566f83"
        }
    ]
}
```
