# Transmission

Social deduction game based on pairwise interaction.

## Dependencies

* [Git](https://git-scm.com/)
* [Go](https://golang.org/)
* [Make](https://www.gnu.org/software/make/)
* [Watchexec](https://github.com/watchexec/watchexec)
* [Yarn](https://yarnpkg.com/)

### For production and deployment

* [Docker](https://www.docker.com/)
* [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli)

## Development

Open two terminals.

```
$ make backend-dev
$ make frontend-dev
```

Game is running at <http://localhost:3455>.

## Production and deployment

Verify outside Docker.

```
$ make backend-build
$ make frontend-build
$ make backend-prod
```

Verify inside Docker.

```
$ make image-build
$ make image-run
```

Deploy.

```
$ make deploy
```

## API

All communication happens via websocket at the single endpoint
`/api/v1/game`. The communication protocol is mostly stateless but not
entirely. Messages are in JSON format.

When you connect to the socket the server will send you

```
{ event: "connected" }
```

followed by

```
{ event: "state", state: { ... } }
```

where the placeholder object is the entire current state of the game,
e.g. whether it's in progress, what players are registered, all the
private data for every player, etc. Clients are given full trust, so
any client can pretend to be any player at any time.

The server will send a new `state` event to every client each time the
state is updated. All clients are expected to update their UI
correspondingly (in this case using the React/Redux paradigm). This is
more or less the extent of the server-to-client messages.

Here is the overall server state object layout (depending on the game
state, some fields may be absent):

```json
{
  "game": {
    "state": "lobby | playing | submission | results",
    "startTime": "2020-11-07T15:29:24Z",
    "stopTime": "2020-11-07T15:39:24Z | null"
  },
  "players": {
    "player-323de6a9-64a0-4974-84f0-2d686f9e5da9": {
      "name": "Radon",
      "color": "#E27D60",
      "node": "node-08424445-0d7c-4c09-8d64-49a113bf0149",
      "role": "headquarters",
      "admin": false
    },
    "player-6b8c5cef-7888-4338-9b3e-5cb814b62408": {
      "name": "Owen",
      "color": "#85DCB",
      "node": "node-a3e0000f-fd5c-42e6-a1fb-060d8c5e5307",
      "role": "trainDepot",
      "admin": false
    },
    "player-2b7d07e2-e692-4092-a888-f63ddcc574ca": {
      "name": "Amit",
      "color": "#E8A87C",
      "node": "node-59ce6684-6597-44de-9be2-fb167b4891d4",
      "role": "doubleAgent",
      "admin": false
    },
    "player-26c36d93-7afc-482a-885a-ba15f3daae6d": {
      "name": "Louise",
      "color": "#41B3A3",
      "node": null,
      "role": null,
      "admin": true
    }
  },
  "graph": {
    "nodes": {
      "node-08424445-0d7c-4c09-8d64-49a113bf0149": [
        "group-f91095f6-d8e8-42f3-b239-79d988437c1b",
        "group-e511762b-9fb9-4921-ae9f-b19b7a49a673"
      ],
      "node-a3e0000f-fd5c-42e6-a1fb-060d8c5e5307": [
        "group-f91095f6-d8e8-42f3-b239-79d988437c1b"
      ],
      "node-59ce6684-6597-44de-9be2-fb167b4891d4": [
        "group-e511762b-9fb9-4921-ae9f-b19b7a49a673"
      ]
    },
    "groups": {
      "group-f91095f6-d8e8-42f3-b239-79d988437c1b": {
        "messages": [
          {
            "timestamp": "2020-11-07T15:29:35Z",
            "sender": "node-08424445-0d7c-4c09-8d64-49a113bf0149",
            "text": "Hello Owen"
          },
          {
            "timestamp": "2020-11-07T15:29:48Z",
            "sender": "node-a3e0000f-fd5c-42e6-a1fb-060d8c5e5307",
            "text": "Hi Radon"
          }
        ]
      },
      "group-e511762b-9fb9-4921-ae9f-b19b7a49a673": {
        "messages": [
          {
            "timestamp": "2020-11-07T15:30:45Z",
            "sender": "node-59ce6684-6597-44de-9be2-fb167b4891d4",
            "text": "Anyone home?"
          }
        ]
      }
    }
  },
  "facts": {
    "possible": {
      "compartment": ["00", "17", "42"],
      "color": ["red", "violet", "red violet", "violet red"],
      "food": ["apple pie", "pecan pie", "pumpkin pie"]
    },
    "true": {
      "compartment": "42",
      "color": "violet red",
      "food": "pecan pie"
    },
    "checked": [
      {
        "field": "color",
        "value": "red violet",
        "accurate": false
      }
    ]
  }
}
```

The client can send messages to update the state.

```json
{
  "event": "updatePlayer",
  "playerID": "2b7d07e2-e692-4092-a888-f63ddcc574ca",
  "playerName": "Amit",
  "playerColor": "#E8A87C",
  "playerNode": "node-59ce6684-6597-44de-9be2-fb167b4891d4",
  "playerAdmin": false
}

{
  "event": "removePlayer",
  "playerID": "6b8c5cef-7888-4338-9b3e-5cb814b62408"
}

{
  "event": "startPregame"
}

{
  "event": "startGame",
  "stopTime": "2020-11-07T15:39:24Z | null"
}

{
  "event": "stopGame"
}

{
  "event": "checkFact",
  "field": "color",
  "value": "red violet"
}

{
  "event": "submitFacts",
  "facts": {
    "compartment": "42",
    "color": "violet red",
    "food": "apple pie"
  }
}

{
  "event": "sendMessage",
  "group": "group-f91095f6-d8e8-42f3-b239-79d988437c1b",
  "sender": "node-08424445-0d7c-4c09-8d64-49a113bf0149",
  "text": "Hello Owen"
}
```
