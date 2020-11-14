<script>
  export let players;
  export let api;
  export let playerID;
  export let toggle;
  let user = getUser(); // TODO: how to recompute when players changes

  // Populates the user object with defaults if no cookie
  // and actual data if the cookie is present
  function getUser() {
    if (players != undefined && players[playerID] != undefined) {
      return {
        joined: true,
        player: players[playerID],
        name: players[playerID].name,
      };
    }
    return {
      joined: false,
      player: null,
      name: "",
    };
  }

  function joinGame(id) {
    if (!user.name) {
      alert("Please enter your name before joining the game.");
      return;
    }
    user.joined = true;
    let message = {
      event: "updatePlayer",
      player: {
        name: user.name,
        id: id,
        color: getColor(id),
        node: null,
        groups: [],
      },
    };
    api.socket.send(JSON.stringify(message));
  }
  function updateName(id) {
    let player = players[id];
    let message = {
      event: "updatePlayer",
      player: {
        ...player,
        name: user.name,
      },
    };
    api.socket.send(JSON.stringify(message));
  }
  function getColor(playerID) {
    if (players && players[playerID] && players[playerID].color) {
      return players[playerID].color;
    }
    return randomColor();
  }
  function startGame() {
    const gameLengthMin = 1;
    let currentTime = new Date();
    let message = {
      event: "startGame",
      stopTime: new Date(currentTime.getTime() + gameLengthMin * 60000),
    };
    api.socket.send(JSON.stringify(message));
  }

  function randomColor() {
    return (
      "#" + ((Math.random() * 0xffffff) << 0).toString(16).padStart(6, "0")
    );
  }
  function removePlayer(id) {
    let message = {
      event: "removePlayer",
      playerID: id,
    };
    api.socket.send(JSON.stringify(message));
    user.joined = false;
  }
</script>

<main>
  Lobby
  {#each Object.entries(players) as [_, { name, color }]}
    <div style="color: {color}">{name}</div>
  {/each}
  {#if !user.joined}
    <div>
      Name:
      <input bind:value={user.name} placeholder="Your name" />
      <button on:click={() => joinGame(playerID)}> Join </button>
    </div>
  {:else}
    <div>
      Update Name:
      <input bind:value={user.name} />
      <button on:click={() => updateName(playerID)}> Submit </button>
      <button on:click={() => removePlayer(playerID)}> Leave Game </button>
    </div>
  {/if}
  <button on:click={startGame}> Start Game </button>
  <button on:click={toggle}> Config </button>
</main>
