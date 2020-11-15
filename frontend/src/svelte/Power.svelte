<script>
 export let role;
 export let state;
 export let api;
 export let playerID;
  import FactChecker from "./roles/FactChecker.svelte";
  import TrainExpert from "./roles/TrainExpert.svelte";
</script>
<main>
    Power:
    {#if role === 'Fact Checker'}
        <FactChecker facts={state.facts} playerID={playerID} api={api} checks={state.players[playerID].checks}/>
    {:else if role === 'Train Expert'}
        <TrainExpert playerID={playerID} players={state.players} api={api}/>
    {:else}
        Not implemented
    {/if}
</main>
