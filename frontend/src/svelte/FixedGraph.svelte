<!-- This is a giant hack. The correct way to do this is to only pass the things that the graph cares about into the graph -->
<script>
  import Graph from "./Graph.svelte";
  export let stateGraph;
  export let api;
  let copy = stateGraph;
</script>

<main>
  <Graph stateGraph={copy} {api} />
</main>
