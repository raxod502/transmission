<script>
  import "./index/global.svelte";

  let state = {
    groups: [
      {
        recipients: [
          "Owen",
        ],
        messages: [
          {
            sender: "Radon",
            text: "Hi Owen",
          },
          {
            sender: "Owen",
            text: "Hello Radon",
          },
        ],
      },
      {
        recipients: [
          "Amit",
        ],
        messages: [
          {
            sender: "Amit",
            text: "Is anyone there?",
          },
        ],
      },
    ],
    facts: {
      real: {
        compartment: {
          name: "Compartment",
          possible: ["00", "17", "42"],
          value: "42",
        },
        color: {
          name: "Color",
          possible: ["red", "violet", "red violet", "violet red"],
          value: "violet red",
        },
        food: {
          name: "Contents",
          possible: ["apple pie", "pecan pie", "pumpkin pie"],
          value: "pecan pie",
        },
      },
    },
  };
</script>

<main>
  <div class="columns is-gapless">
    <div class="column is-three-quarters">
      <div class="rows" style="height: 100vh">
        <div class="row" style="height: 20%">
          <div class="columns">
            <div class="column">
              Name and role
            </div>
            <div class="column">
              Timer
            </div>
          </div>
        </div>
        <div class="row" style="height: 80%">
          <div class="columns">
            {#each state.groups as { recipients, messages }}
              <div class="column">
                <div class="rows">
                  <div class="row">
                    Conversation with { recipients.join(", ") }
                  </div>
                  <div class="row">
                    {#each messages as { sender, text }}
                      <p>{sender}: {text}</p>
                    {/each}
                  </div>
                </div>
              </div>
            {/each}
          </div>
        </div>
      </div>
    </div>
    <div class="column">
      <div class="rows" style="height: 100vh">
        <div class="row" style="height: 25%">
          Network
        </div>
        <div class="row" style="height: 25%">
          <p>Facts</p>
          {#each Object.entries(state.facts.real) as [id, { name, possible, value }]}
            <p>
              <label for="fact-dropdown-{id}">{name}</label>
              <select name="fact-dropdown-{id}" value={value}>
                {#each possible as value}
                  <option value="{value}">{value}</option>
                {/each}
              </select>
            </p>
          {/each}
        </div>
        <div class="row" style="height: 25%">
          <p>Notes</p>
          <textarea></textarea>
        </div>
        <div class="row" style="height: 25%">
          <p>Power</p>
        </div>
      </div>
    </div>
  </div>
</main>
