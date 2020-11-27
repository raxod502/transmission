<script>
  export let stateGraph;

  import sigma from "sigma";
  import { onMount } from "svelte";

  let container;
  let sigmaGraph;

  function updateGraph(s, stateGraph) {
    if (s === undefined) return;
    for (const node of s.graph.nodes()) {
      s.graph.dropNode(node.id);
    }
    const edges = new Map();
    const entries = Object.entries(stateGraph.nodes);
    for (const [idx, [nodeID, { groups: groupIDs }]] of entries.entries()) {
      const x = Math.cos((idx / entries.length) * 2 * Math.PI);
      const y = Math.sin((idx / entries.length) * 2 * Math.PI);
      s.graph.addNode({
        id: nodeID,
        label: stateGraph.nodes[nodeID].name,
        x: x,
        y: y,
        size: 1,
        color: stateGraph.nodes[nodeID].color,
      });
      for (const groupID of groupIDs) {
        if (groupID === "group-baddies") {
          continue;
        }
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
        continue;
      }
      const [source, target] = Array.from(groupNodeIDs);
      s.graph.addEdge({
        id: groupID,
        source,
        target,
        color: "#000",
      });
    }
    s.refresh();
  }

  $: updateGraph(sigmaGraph, stateGraph);

  onMount(() => {
    sigmaGraph = new sigma({
      container,
      settings: {
        mouseEnabled: false,
      },
    });
    window.sigma = sigmaGraph;
  });
</script>

<div bind:this={container} style="height: 200px" />
