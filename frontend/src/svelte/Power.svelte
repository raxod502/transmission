<script>
 export let role;
 export let state;
 export let api;
 export let playerID;
 export let powerUses;
 export let selectedFacts;
  import FactChecker from "./roles/FactChecker.svelte";
  import TrainExpert from "./roles/TrainExpert.svelte";
  import TrainDepot from "./roles/TrainDepot.svelte";
  import Headquarters from "./roles/Headquarters.svelte";
</script>
<main>
    Power:
    {#if role === 'Fact Checker'}
        <FactChecker facts={state.facts} playerID={playerID} api={api} checks={state.players[playerID].checks} powerUses={powerUses}/>
    {:else if role === 'Train Expert'}
        <TrainExpert playerID={playerID} players={state.players} api={api} nodes={state.graph.nodes} powerUses={powerUses}/>
    {:else if role === 'Train Depot'}
        <TrainDepot selectedFacts={selectedFacts} api={api}/>
    {:else if role === 'Double Agent'}
        <TrainExpert playerID={playerID} players={state.players} api={api} role={"Double Agent"} nodes={state.graph.nodes} powerUses={powerUses}/>
    {:else if role === 'Headquarters'}
        <Headquarters api={api} facts={state.facts} playerID={playerID}/>
    {:else}
        Not implemented
    {/if}
</main>
