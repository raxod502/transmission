<script>
  import "./index/global.svelte";

  import { API } from "./index/api.js";

  let state = {
    game: {
      state: "loading",
    },
  };

  const api = new API({
    onStateUpdate: (newState) => {
      console.log("Received state update:", newState);
      state = newState;
    },
  });
  api.connect();
</script>

<main>
  {#if state.game.state === 'loading'}
    Loading...
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
