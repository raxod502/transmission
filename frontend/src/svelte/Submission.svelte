<script>
  export let player;
  export let api;
  export let toggle;
  export let facts;
  let submission = {};
  function sendSubmission() {
    let message = {
      event: "submitFacts",
      submission: submission,
    };
    api.send(message);
  }
</script>

<main>
  {#if player.role == 'Train Depot'}
    Submit answers:
    {#each Object.entries(facts.real) as [name, fact]}
      <div>
        {name}:
        <select bind:value={submission[name]}>
          {#each fact.possible as value}
            <option {value}>{value}</option>
          {/each}
        </select>
      </div>
    {/each}
    <button on:click={sendSubmission}> Submit </button>
  {:else}Waiting for answers to be submitted{/if}
  <div><button on:click={toggle}> Config </button></div>
</main>
