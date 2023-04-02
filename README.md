# go_reservation_api

## What is this?

This project is a learning experience in Go programing, most especifcally a REST API.

I do have previous experience with Go but mostly in small scripts to run in old dedicated Linux servers.

## Will this become something?

Yes. A simple API with basic user functionality and a basic sports facility reservation functionality. Honestly, it could very much change in the future if I feel more creative with it (or it becomes a fork).

## How can I run it?

First, get yourself a Postgress Database. I prefer a local one, so use `docker-compose -f docker-compose-dev.yml up`. After it's done, just run `make run`.

If for some reason you need to change credentials edit the Makefile `DATABASE_URL` env.

It has a basic migration functionality (to be refactored.... one day), so just running it will do everything and set you up correctly.
