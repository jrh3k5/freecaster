# Freecaster

Freecaster is a Farcaster bot that publishes posts about free games available to it. It is powered by the [FreeStuff API](https://docs.freestuffbot.xyz/) and hosted on [Vercel](https://vercel.com).

## Running Locally

This project supports running locally with Docker. First, you need to create a `.env` file. Do the following:

```
cp .env.sample .env
```

Then fill out the configuration properties. Following that, you can run the application by executing:

```
docker compose up --build
```

In another window, execute:

```
curl -v POST "http://localhost:8080/api/index" \
  -d '{ "event": "free_games", "secret": "wdaji29dJadj91jAjd9a92eDak2", "data": [ 142312, 499128 ] }'
```

## Configuring

This application needs the following configuration parameters as environmental variables:

* `FREESTUFF_WEBHOOK_SECRET`: the pre-configured secret in the FreeStuff API used to help verify that the request originates from the FreeStuff API

### Optional Configuration

* `LOG_LEVEL`: by default, the application logs at warn level. This environmental variable can be set to a value correspond to a Zap log level to override that default.