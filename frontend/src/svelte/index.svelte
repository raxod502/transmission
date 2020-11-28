<script>
  import "./index/global.svelte";

  import { API } from "./index/api.js";
  import { v4 as uuidv4 } from "uuid";
  import FixedGraph from "./FixedGraph.svelte";
  import Timer from "./Timer.svelte";
  import Lobby from "./Lobby.svelte";
  import Config from "./Config.svelte";
  import Submission from "./Submission.svelte";
  import Results from "./Results.svelte";
  import Power from "./Power.svelte";

  let state = {
    game: {
      state: "loading",
    },
  };
  let selectedFacts = {};
  let config = false;
  let playerID = getPlayerID();
  let composedMessages = {};

  const api = new API({
    onStateUpdate: (newState) => {
      console.log("Received state update:", newState);
      state = newState;
      state.players = newState.players;
      window.state = state;
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
    api.send(message);
  }
  function checkForSend(groupID) {
    if (event.code == "Enter" && composedMessages[groupID] !== "") {
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
    api.send(message);
  }
  function getInterlocutorName(playerID, groupID) {
    const curNodeID = state.players[playerID].node;
    for (const [id, { groups }] of Object.entries(state.graph.nodes)) {
      if (id !== curNodeID && groups.includes(groupID)) {
        return state.graph.nodes[id].name;
      }
    }
    return "Unknown Interlocutor";
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
                <p>Name: {state.players[playerID].name}</p>
                <p>Role: {state.players[playerID].role}</p>
                {#if state.players[playerID].knownFacts}
                  {#each Object.entries(state.players[playerID].knownFacts) as [name, _]}
                    <div>
                      You know that the
                      {name}
                      is
                      {state.facts.real[name].value}
                    </div>
                  {/each}
                {/if}
              </div>
              <div class="column">
                <p class="has-text-weight-bold">Timer:</p>
                <Timer
                  startTime={state.game.startTime}
                  endTime={state.game.stopTime}
                  {api} />
              </div>
            </div>
          </div>
          <div class="row" style="height: 80%">
            <div class="columns is-gapless mx-2" style="height: 100%">
              {#if Object.keys(state.graph.groups).length === 0}
                <p>No conversations</p>
              {:else if state.players[playerID].node === ""}
                <p>You weren't assigned to a node :(</p>
              {:else}
                {#each state.graph.nodes[state.players[playerID].node].groups as groupID}
                  {#if state.graph.groups[groupID]}
                    <div class="column" style="height: 100%">
                      <div class="rows" style="height: 100%; display: flex">
                        <div class="row mx-3" style="width: 100%; height: 100%">
                          <div
                            style="overflow-y: auto; height: calc(100% - 60px); display: flex">
                            <div style="margin-top: auto">
                              {#each state.graph.groups[groupID].messages.slice(-20) as { sender, text }}
                                <p
                                  style="color: {state.graph.nodes[sender].color}"
                                  class="mb-3">
                                  {state.graph.nodes[sender].name}:
                                  {text}
                                </p>
                              {/each}
                            </div>
                          </div>
                          <div style="height: 60px">
                            <input
                              type="text"
                              class="input"
                              on:keyup|preventDefault={() => checkForSend(groupID)}
                              bind:value={composedMessages[groupID]}
                              style="width: 100%"
                              placeholder="Message to {getInterlocutorName(playerID, groupID)}" />
                          </div>
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
            <p class="has-text-weight-bold">Network</p>
            <FixedGraph stateGraph={state.graph} {api} />
          </div>
          <div class="row" style="height: 25%">
            <p class="has-text-weight-bold">Facts</p>
            {#each Object.entries(state.facts.real) as [name, { possible, value }]}
              <p>
                <label for="fact-dropdown-{name}">{name}</label>
                <select
                  name="fact-dropdown-{name}"
                  bind:value={selectedFacts[name]}>
                  {#each possible as value}
                    <option {value}>{value}</option>
                  {/each}
                </select>
              </p>
            {/each}
          </div>
          <div class="row" style="height: 25%">
            <p class="has-text-weight-bold">Notes</p>
            <textarea />
          </div>
          <div class="row" style="height: 25%">
            <Power
              role={state.players[playerID].role}
              powerUses={state.players[playerID].powerUses}
              {state}
              {api}
              {playerID}
              {selectedFacts} />
            <p class="has-text-weight-bold">Admin:</p>
            <button class="button is-danger" on:click={toggleConfig}>
              Config Panel
            </button>
            <button class="button is-link" on:click={goToLobby}>
              Return To Lobby
            </button>
          </div>
        </div>
      </div>
    </div>
  {:else if state.game.state === 'results'}
    <Results facts={state.facts} players={state.players} />
    <button class="button is-link" on:click={goToLobby}>
      Return To Lobby
    </button>
  {/if}
</main>
