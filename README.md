# Welcome to Revel

A high-productivity web framework for the [Go language](http://www.golang.org/).

### Start the web server:

   revel run revel-rest

### Go to http://localhost:9000/ and you'll see:

    "It works"

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites

## API Document:

### GET Trains:
curl --location --request GET 'http://localhost:9000/v1/trains'

### GET Train by ID:
curl --location --request GET 'http://localhost:9000/v1/trains/1'

### CREATE Train:
curl --location --request POST 'http://localhost:9000/v1/trains' \
--header 'Content-Type: application/json' \
--data-raw '{
  "driver_name": "vin",
  "operating_status": true
}'

### UPDATE Train:
curl --location --request PATCH 'http://localhost:9000/v1/trains/1' \
--header 'Content-Type: application/json'
--data-raw '{
  "driver_name": "Live",
  "operating_status": false
}'

### DELETE train id:
curl --location --request DELETE 'http://localhost:9000/v1/trains/1'

## Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

