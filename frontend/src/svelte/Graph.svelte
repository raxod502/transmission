<script>
  export let stateGraph;

  import sigma from "sigma";
  import { onMount } from "svelte";

  let container;
  let sigmaGraph;

  function updateGraph(s, stateGraph) {
    if (s === undefined) return;
    const edges = new Map();
    for (const [nodeID, { groups: groupIDs }] of Object.entries(
      stateGraph.nodes
    )) {
      console.log("add node", nodeID);
      s.graph.addNode({
        id: nodeID,
        label: nodeID,
      });
      for (const groupID of groupIDs) {
        let groupNodeIDs = edges.get(groupID);
        if (!groupNodeIDs) {
          groupNodeIDs = new Set();
        }
        groupNodeIDs.add(nodeID);
        edges.set(groupID, groupNodeIDs);
      }
    }
    for (const [groupID, groupNodeIDs] of edges.entries()) {
      if (groupNodeIDs.size !== 2) {
        console.log(
          `Graph rendering: skipping hyper-edge of size ${groupNodeIDs.size}:`,
          groupID
        );
        continue;
      }
      const [source, target] = Array.from(groupNodeIDs);
      console.log("add edge", groupID, `[${source} -> ${target}]`);
      s.graph.addEdge({
        id: groupID,
        source,
        target,
      });
    }
    s.refresh();
  }

  $: updateGraph(sigmaGraph, stateGraph);

  onMount(() => {
    console.log("container:", container);
    sigmaGraph = new sigma({
      container,
    });
    window.sigma = sigmaGraph;
  });
</script>

<div bind:this={container} />
