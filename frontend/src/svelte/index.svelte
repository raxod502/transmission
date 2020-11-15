<script>
  import "./index/global.svelte";

  import { API } from "./index/api.js";
  import { v4 as uuidv4 } from "uuid";
  import Graph from "./Graph.svelte";
  import Timer from "./Timer.svelte";
  import Lobby from "./Lobby.svelte";
  import Config from "./Config.svelte";
  import Submission from "./Submission.svelte";
  import Power from "./Power.svelte";

  let state = {
    game: {
      state: "loading",
    },
  };
  let config = false;
  let playerID = getPlayerID();
  let composedMessages = {};

  const api = new API({
    onStateUpdate: (newState) => {
      console.log("Received state update:", newState);
      state = newState;
      state.players = newState.players;
    },
  });
  api.connect();

  function toggleConfig() {
    config = !config;
  }
  function getPlayerIDFromCookie() {
    let currentID = document.cookie
      .split("; ")
      .find((row) => row.startsWith("playerID"));
    if (currentID) {
      return currentID.split("=")[1];
    }
    return null;
  }

  // Reads playerID from the cookie. Sets a random one if it doesn't exist
  function getPlayerID() {
    let currentID = getPlayerIDFromCookie();
    if (currentID) {
      return currentID;
    }
    let newUUID = "player-" + uuidv4();
    document.cookie =
      "playerID=" + newUUID + "; expires=Fri, 31 Dec 9999 23:59:59 GMT";
    return newUUID;
  }
  function goToLobby() {
    let message = {
      event: "startPregame",
    };
    api.socket.send(JSON.stringify(message));
  }
  function checkForSend(groupID) {
    if (event.code == "Enter") {
      console.log("sending", composedMessages[groupID], "to", groupID);
      sendMessage(
        groupID,
        state.players[playerID].node,
        composedMessages[groupID]
      );
      composedMessages[groupID] = "";
    }
  }
  function sendMessage(groupID, sender, text) {
    let message = {
      event: "sendMessage",
      groupID: groupID,
      sender: sender,
      text: text,
    };
    api.socket.send(JSON.stringify(message));
  }
</script>

<main>
  {#if config}
    <Config
      players={state.players}
      availableRoles={state.possibleRoles}
      {toggleConfig}
      {api}
      graph={state.graph}
      facts={state.facts}
      possibleFacts={state.possibleFacts} />
  {:else if state.game.state === 'loading'}
    Loading...
  {:else if state.game.state === 'lobby'}
    <Lobby players={state.players} {api} {playerID} toggle={toggleConfig} />
  {:else if state.game.state === 'submission'}
    <Submission
      player={state.players[playerID]}
      {api}
      toggle={toggleConfig}
      facts={state.facts} />
  {:else if state.game.state === 'playing'}
    <div class="columns is-gapless">
      <div class="column is-three-quarters">
        <div class="rows" style="height: 100vh">
          <div class="row" style="height: 20%">
            <div class="columns is-gapless">
              <div class="column">
                Name:
                {state.players[playerID].name}
                and role:
                {state.players[playerID].role}
              </div>
              <div class="column">
                Timer:
                <Timer
                  startTime={state.game.startTime}
                  endTime={state.game.stopTime}
                  {api} />
              </div>
            </div>
          </div>
          <div class="row" style="height: 80%">
            <div class="columns is-gapless">
              {#if Object.keys(state.graph.groups).length === 0}
                <p>No conversations</p>
              {:else}
                {#each state.graph.nodes[state.players[playerID].node].groups as groupID}
                  {#if state.graph.groups[groupID]}
                    <div class="column">
                      <div class="rows">
                        <div class="row">Conversation with: TODO</div>
                        <div class="row">
                          {#each state.graph.groups[groupID].messages as { sender, text }}
                            <p style="color: {state.graph.nodes[sender].color}">
                              {state.graph.nodes[sender].name}:
                              {text}
                            </p>
                          {/each}
                          <input
                            type="text"
                            on:keyup|preventDefault={() => checkForSend(groupID)}
                            bind:value={composedMessages[groupID]} />
                        </div>
                      </div>
                    </div>
                  {/if}
                {/each}
              {/if}
            </div>
          </div>
        </div>
      </div>
      <div class="column">
        <div class="rows" style="height: 100vh">
          <div class="row" style="height: 25%">
            <p>Network</p>
            <Graph stateGraph={state.graph} />
          </div>
          <div class="row" style="height: 25%">
            <p>Facts</p>
            {#each Object.entries(state.facts.real) as [name, { possible, value }]}
              <p>
                <label for="fact-dropdown-{name}">{name}</label>
                <select name="fact-dropdown-{name}" {value}>
                  {#each possible as value}
                    <option {value}>{value}</option>
                  {/each}
                </select>
              </p>
            {/each}
          </div>
          <div class="row" style="height: 25%">
            <p>Notes</p>
            <textarea />
          </div>
          <div class="row" style="height: 25%">
            <Power role={state.players[playerID].role} state={state} api={api} playerID={playerID}/>
            <button on:click={toggleConfig}> Config Panel </button>
            <button on:click={goToLobby}> Return To Lobby </button>
          </div>
        </div>
      </div>
    </div>
  {:else if state.game.state === 'results'}
    TODO: implement results page
    <button on:click={goToLobby}> Return To Lobby </button>
  {/if}
</main>
