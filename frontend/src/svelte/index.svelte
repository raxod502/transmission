<script>
  import "./index/global.svelte";

  import { API } from "./index/api.js";
  import { v4 as uuidv4 } from 'uuid';

  let state = {
    game: {
      state: "loading",
    },
  };
  let user = getUser()

  const api = new API({
    onStateUpdate: (newState) => {
      console.log("Received state update:", newState);
      state = newState;
      user = getUser(); //TODO: figure out how to make this reactive the nice way

    },
  });
  api.connect();

 function getPlayerIDFromCookie(){
     let currentID = document.cookie
                             .split('; ')
                             .find(row => row.startsWith("playerID"))
     if (currentID){
        return currentID.split('=')[1];
     }
     return null
 }

 // Reads playerID from the cookie. Sets a random one if it doesn't exist
 function getPlayerID() {
     let currentID = getPlayerIDFromCookie()
     if (currentID) {
         return currentID
     }
     let newUUID = "player-" + uuidv4()
     document.cookie = "playerID=" + newUUID + "; expires=Fri, 31 Dec 9999 23:59:59 GMT"
     return newUUID
 }

 function startGame(){
     const gameLengthMin = 10;
     let currentTime = new Date();
     let message = {
         event: "startGame",
         stopTime: new Date(currentTime.getTime() + gameLengthMin * 60000),
     };
    api.socket.send(JSON.stringify(message));
 }

 // Populates the user object with defaults if no cookie
 // and actual data if the cookie is present
 function getUser(){
     let playerID = getPlayerID();
     if (state.players != undefined && state.players[playerID]!= undefined){
         return {joined: true, name: state.players[playerID].name}
     }
     return {joined: false, name: "Set your name here"}
 }

 function getColor(playerID){
     if (state.players && state.players[playerID] && state.players[playerID].color){
         return state.players[playerID].color
     }
     return randomColor()
 }

 function randomColor(){
    return '#'+(Math.random() * 0xFFFFFF << 0).toString(16).padStart(6, '0')
 }

 function updateName(){
     let id = getPlayerID();
     let player = state.players[id]
     let message = {
         event: "updatePlayer"
         player: {
             ...player,
             name: user.name
         }
     };
     api.socket.send(JSON.stringify(message));
 }

 function joinGame(){
     user.joined = true;
     let id = getPlayerID();
     let message = {
         event: "updatePlayer" ,
         player: {
             name: user.name,
             id: id,
             color: getColor(id),
             node: null,
             groups: [],
         }
     };
     api.socket.send(JSON.stringify(message));
 }
</script>

<main>
  {#if state.game.state === 'loading'}
    Loading...
  {:else if state.game.state === 'lobby'}
      Lobby
      {#each Object.entries(state.players) as [_, { name, color }]}
          <div style="color: {color}"> {name}</div>
      {/each}
      {#if !user.joined}
          <div>
            Name:
            <input bind:value={user.name}>
            <button on:click={joinGame}> Join </button>
          </div>
      {:else}
          <div>
            Update Name:
            <input bind:value={user.name}>
            <button on:click={updateName}> Submit </button>
          </div>
      {/if}
      <button on:click={startGame}> Start Game </button>
  {:else}
    <div class="columns is-gapless">
      <div class="column is-three-quarters">
        <div class="rows" style="height: 100vh">
          <div class="row" style="height: 20%">
            <div class="columns is-gapless">
              <div class="column">Name and role</div>
              <div class="column">Timer</div>
            </div>
          </div>
          <div class="row" style="height: 80%">
            <div class="columns is-gapless">
              {#if Object.keys(state.graph.groups).length === 0}
                <p>No conversations</p>
              {:else}
                {#each Object.entries(state.graph.groups) as [_, { recipients, messages }]}
                  <div class="column">
                    <div class="rows">
                      <div class="row">
                        Conversation with
                        {recipients.join(', ')}
                      </div>
                      <div class="row">
                        {#each messages as { sender, text }}
                          <p>{sender}: {text}</p>
                        {/each}
                      </div>
                    </div>
                  </div>
                {/each}
              {/if}
            </div>
          </div>
        </div>
      </div>
      <div class="column">
        <div class="rows" style="height: 100vh">
          <div class="row" style="height: 25%">Network</div>
          <div class="row" style="height: 25%">
            <p>Facts</p>
            {#each Object.entries(state.facts.real) as [id, { name, possible, value }]}
              <p>
                <label for="fact-dropdown-{id}">{name}</label>
                <select name="fact-dropdown-{id}" {value}>
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
            <p>Power</p>
          </div>
        </div>
      </div>
    </div>
  {/if}
</main>
