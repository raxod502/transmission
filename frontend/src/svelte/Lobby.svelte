<script>
  export let players;
  export let api;
  export let playerID;
  export let toggle;
  let gameLength = 10;
  let user = getUser(); // TODO: how to recompute when players changes

  // https://sashamaps.net/docs/resources/20-colors/
  const distinctColors = [
    "#e6194B", // Red
    "#3cb44b", // Green
    "#4363d8", // Blue
    "#f58231", // Orange
    "#911eb4", // Purple
    "#f032e6", // Magenta
    "#469990", // Teal
    "#9A6324", // Brown
    "#800000", // Maroon
    "#808000", // Olive
    "#000075", // Navy
    "#a9a9a9", // Grey
  ];

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
      },
    };
    api.send(message);
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
    api.send(message);
  }
  function getColor(playerID) {
    if (players && players[playerID] && players[playerID].color) {
      return players[playerID].color;
    }
    return randomColor();
  }
  function startGame() {
    const gameLengthMin = gameLength;
    let currentTime = new Date();
    let message = {
      event: "startGame",
      stopTime: new Date(currentTime.getTime() + gameLengthMin * 60000),
    };
    api.send(message);
  }

  function randomColor() {
    const usedColors = new Set();
    for (const [_, { color }] of Object.entries(players)) {
      usedColors.add(color);
    }
    for (const color of distinctColors) {
      if (!usedColors.has(color)) {
        return color;
      }
    }
    // If all distinct colors have been used, just pick something
    // random.
    return (
      "#" + ((Math.random() * 0xffffff) << 0).toString(16).padStart(6, "0")
    );
  }
  function removePlayer(id) {
    let message = {
      event: "removePlayer",
      playerID: id,
    };
    api.send(message);
    user.joined = false;
  }
</script>

<main
  style="min-height: 100vh; display: flex; justify-content:
             center; align-items: center">
  <div>
    <div class="has-text-centered">
      <b style="text-decoration: underline"> Players in the lobby </b>
      {#each Object.values(players).sort(({ name: n1 }, { name: n2 }) =>
        n1.localeCompare(n2)
      ) as { name, color }}
        <div style="color: {color}">{name}</div>
      {/each}
    </div>
    {#if !user.joined}
      <div class="my-5">
        Name:
        <form class="field is-grouped">
          <p class="control">
            <input
              bind:value={user.name}
              class="input"
              placeholder="Your name"
              autofocus />
          </p>
          <p class="control">
            <button
              type="submit"
              class="button is-link"
              on:click={() => joinGame(playerID)}>
              Join
            </button>
          </p>
        </form>
      </div>
    {:else}
      <div class="my-5">
        Update name:
        <form class="field is-grouped">
          <p class="control"><input bind:value={user.name} class="input" /></p>
          <p class="control">
            <button
              type="submit"
              class="button is-link"
              on:click={() => updateName(playerID)}>Update name</button>
          </p>
          <p class="control">
            <button
              type="button"
              class="button is-danger"
              on:click={() => removePlayer(playerID)}>Leave Game</button>
          </p>
        </form>
      </div>
      <div class="my-5">
        Start game:
        <form class="field is-grouped">
          <p class="control">
            <input
              class="input"
              type="number"
              bind:value={gameLength}
              min="0"
              max="480" />
          </p>
          <p class="control">
            <button type="submit" class="button is-success" on:click={startGame}>
              Start Game
            </button>
          </p>
        </form>
      </div>
    {/if}
    <button on:click={toggle} class="my-5 is-pulled-right button is-light">
      Admin panel
    </button>
  </div>
</main>
