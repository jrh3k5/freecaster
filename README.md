# Freecaster

Freecaster is a Farcaster bot that publishes posts about free games available to it. It is powered by the [FreeStuff API](https://docs.freestuffbot.xyz/) and hosted on [Vercel](https://vercel.com).

## Running Locally

This project supports running locally with Docker. To run it, execute:

```
docker compose up --build
```

In another window, execute:

```
curl -v POST "http://localhost:8080/api/index" \
  -d '{ "event": "free_games", "secret": "wdaji29dJadj91jAjd9a92eDak2", "data": [ 142312, 499128 ] }'
```