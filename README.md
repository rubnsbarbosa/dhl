# DHL Tracker CLI

```bash                   
 ______________  _______
 ___  __ \__  / / /__  /
 __  / / /_  /_/ /__  /
 _  /_/ /_  __  / _  /___
 /_____/ /_/ /_/  /_____/
 
 ~ DHL Tracker CLI 1.0 ~
```

## Prerequisites 

You must have an eligible **Tracking Code** for the shipment and an **API Key** subscription which is specified in the request header.

## Setup

DHL Tracker CLI requires an API key. To run it:

1. Obtain your API key from [Developer DHL](https://developer.dhl.com)
2. Set Environment Variable

```bash
export API_KEY="YOUR_DHL_API_KEY"
```

Replace `"YOUR_DHL_API_KEY"` with your API key value

## How to run

```bash
go build main.go tracking.go
./main -tracking RR000000001DE
```

or with Docker

```bash
docker pull rubnsbarbosa/dhl
docker run -e API_KEY=YOUR_DHL_API_KEY dhl -tracking RR000000001DE
```

## Reference

* Find more details about the docker image on [Docker Hub - rubnsbarbosa/dhl](https://hub.docker.com/r/rubnsbarbosa/dhl)
* Find the steps on how to obtain your API key from [Developer DHL Shipment Tracking Unified API](https://developer.dhl.com/api-reference/shipment-tracking)

## License<a id="license"></a>

This repository is licensed under the [MIT License](https://github.com/rubnsbarbosa/dhl/blob/main/LICENSE)
