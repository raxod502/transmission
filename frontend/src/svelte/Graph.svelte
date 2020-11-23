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
    for (const [nodeID, { groups: groupIDs }] of Object.entries(
      stateGraph.nodes
    )) {
      // console.log("add node", nodeID);
      // console.log(s.graph.nodes());
      var x;
      var y;
      if (nodeID.length < 10) {
        console.log(nodeID);
        let num = parseInt(nodeID.slice(5), 10);
        x = Math.cos((num / 6) * 2 * Math.PI);
        y = Math.sin((num / 6) * 2 * Math.PI);
      } else {
        x = Math.random();
        y = Math.random();
      }
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
        // console.log(
        //   `Graph rendering: skipping hyper-edge of size ${groupNodeIDs.size}:`,
        //   groupID
        // );
        continue;
      }
      const [source, target] = Array.from(groupNodeIDs);
      // console.log("add edge", groupID, `[${source} -> ${target}]`);
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
    // console.log("container:", container);
    sigmaGraph = new sigma({
      container,
    });
    window.sigma = sigmaGraph;
  });
</script>

<div bind:this={container} style="height: 200px" />
