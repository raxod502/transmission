# Transmission

Social deduction game based on pairwise interaction.

## Dependencies

* [Git](https://git-scm.com/)
* [Go](https://golang.org/)
* [Make](https://www.gnu.org/software/make/)
* [Watchexec](https://github.com/watchexec/watchexec)
* [Yarn](https://yarnpkg.com/)

To install the frontend dependencies, run:

```
$ yarn install
```

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
      "id": "player-323de6a9-64a0-4974-84f0-2d686f9e5da9"
      "color": "#E27D60",
      "node": "node-08424445-0d7c-4c09-8d64-49a113bf0149",
      "role": "HQ",
      "checks": [
        {
          "name": "color",
          "guessedValue": "green"
          "correct": false
        }
      ]
    },
    "player-6b8c5cef-7888-4338-9b3e-5cb814b62408": {
      "name": "Owen",
      "id": "player-6b8c5cef-7888-4338-9b3e-5cb814b62408",
      "color": "#85DCB",
      "node": "node-a3e0000f-fd5c-42e6-a1fb-060d8c5e5307",
      "role": "TD",
      "checks": []
    },
    "player-2b7d07e2-e692-4092-a888-f63ddcc574ca": {
      "name": "Amit",
      "id": "player-2b7d07e2-e692-4092-a888-f63ddcc574ca",
      "color": "#E8A87C",
      "node": "node-59ce6684-6597-44de-9be2-fb167b4891d4",
      "role": "baddie",
      "checks": []
    }
  },
  "graph": {
    "nodes": {
      "node-08424445-0d7c-4c09-8d64-49a113bf0149": {
        "id": "node-08424445-0d7c-4c09-8d64-49a113bf0149",
        "player": "player-6b8c5cef-7888-4338-9b3e-5cb814b62408",
        "name": "Owen",
        "color": "#85DCB",
        "groups" : [
          "group-f91095f6-d8e8-42f3-b239-79d988437c1b",
          "group-e511762b-9fb9-4921-ae9f-b19b7a49a673"
        ]
      },
      "node-a3e0000f-fd5c-42e6-a1fb-060d8c5e5307": {
        "id": "node-a3e0000f-fd5c-42e6-a1fb-060d8c5e5307",
        "player": "player-2b7d07e2-e692-4092-a888-f63ddcc574ca",
        "name": "Amit",
        "color": "#E8A87C",
        "groups":[
            "group-f91095f6-d8e8-42f3-b239-79d988437c1b"
          ]
      },
      "node-59ce6684-6597-44de-9be2-fb167b4891d4": {
        "id": "node-59ce6684-6597-44de-9be2-fb167b4891d4",
        "player": "player-26c36d93-7afc-482a-885a-ba15f3daae6d",
        "name": "Louise",
        "color": "#41B3A3",
        "groups": [
            "group-e511762b-9fb9-4921-ae9f-b19b7a49a673"
          ]
      }
    },
    "groups": {
      "group-f91095f6-d8e8-42f3-b239-79d988437c1b": {
      "id": "group-f91095f6-d8e8-42f3-b239-79d988437c1b",
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
      "id": "group-e511762b-9fb9-4921-ae9f-b19b7a49a673",
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
    "real": {
      "compartment": {
        "possible": ["00", "17", "42"],
        "value": "42"
      },
      "color": {
        "possible": ["red", "violet", "red violet", "violet red"],
        "value": "violet red"
      },
      "food": {
        "possible": ["apple pie", "pecan pie", "pumpkin pie"],
        "value": "pecan pie"
      }
    }
  }
}
```

The client can send messages to update the state.

```json
{
  "event": "updatePlayer",
  "player": {
    "name": "Amit",
    "node": "node-59ce6684-6597-44de-9be2-fb167b4891d4",
    "id": "2b7d07e2-e692-4092-a888-f63ddcc574ca",
    "role": "HQ",
    "color": "#E8A87C",
    "admin": false,
    "checks": []
  }
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
  "stopTime": "2020-11-07T15:39:24Z" | null
}

{
  "event": "stopGame"
}

{
  "event": "checkFact",
  "playerID": "6b8c5cef-7888-4338-9b3e-5cb814b62408",
  "field": "color",
  "value": "red violet"
}

{
  "event": "submitFacts",
  "submission": {
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

{
"event": "updateNode",
  "node": {
        "id": "node-08424445-0d7c-4c09-8d64-49a113bf0149",
        "player": "player-6b8c5cef-7888-4338-9b3e-5cb814b62408",
        "name": "Owen",
        "color": "#85DCB",
        "groups" : [
          "group-f91095f6-d8e8-42f3-b239-79d988437c1b",
          "group-e511762b-9fb9-4921-ae9f-b19b7a49a673"
        ]
      }
}

{
  "event": "updateGroup",
  "group": {
        "id": "group-f91095f6-d8e8-42f3-b239-79d988437c1b",
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
      }
}
{
  "event": "updateRealFactPossibilities",
  "factName": "color",
  "possibleValues": ["red", "violet", "red violet", "violet red"]
}
```
