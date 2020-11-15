<script>
  import { v4 as uuidv4 } from "uuid";
  import Graph from "./Graph.svelte";

  export let players;
  export let availableRoles;
  export let toggleConfig;
  export let api;
  export let graph;
  export let possibleFacts;
  export let facts;
  let factsIncluded = {};

  function updatePlayer(id) {
    let player = players[id];
    console.log(player);
    let playerMessage = {
      event: "updatePlayer",
      player: player,
    };
    let nodeMessage = {
      event: "updateNode",
      node: {
        ...graph.nodes[player.node],
        player: player.id,
        name: player.name,
        color: player.color,
      },
    };
    api.socket.send(JSON.stringify(playerMessage));
    api.socket.send(JSON.stringify(nodeMessage));
  }
  function addNode() {
    let message = {
      event: "updateNode",
      node: {
        id: "node-" + uuidv4(),
        groups: [],
      },
    };
    api.socket.send(JSON.stringify(message));
  }
  function removeNode(id) {
    let message = {
      event: "removeNode",
      nodeID: id,
    };
    api.socket.send(JSON.stringify(message));
  }
  function addGroup() {
    let message = {
      event: "updateGroup",
      group: {
        id: "group-" + uuidv4(),
        messages: [],
      },
    };
    api.socket.send(JSON.stringify(message));
  }
  function removeGroup(id) {
    let message = {
      event: "removeGroup",
      groupID: id,
    };
    api.socket.send(JSON.stringify(message));
  }
  let groupsToAdd = {};
  function associateGroup(nodeID) {
    let groupID = groupsToAdd[nodeID];
    let node = graph.nodes[nodeID];
    let message = {
      event: "updateNode",
      node: {
        ...node,
        groups: [...node.groups, groupID],
      },
    };
    api.socket.send(JSON.stringify(message));
  }
  function disassociateGroup(node, groupID) {
    console.log(groupID);
    console.log(node.groups);
    let newGroups = node.groups.filter((g) => g !== groupID);
    console.log(node.groups);
    let message = {
      event: "updateNode",
      node: {
        ...node,
        groups: newGroups,
      },
    };
    api.socket.send(JSON.stringify(message));
  }
  function selectActiveFacts() {
    let facts = {};
    for (const [name, fact] of Object.entries(possibleFacts)) {
      if (factsIncluded[name]) {
        facts[name] = fact;
      }
    }
    setRealFacts(facts);
  }
  function setRealFacts(facts) {
    let message = {
      event: "setRealFacts",
      facts: facts,
    };
    api.socket.send(JSON.stringify(message));
  }
</script>

<main>
  <p>Config</p>
  <p>Players:</p>
  {#each Object.entries(players) as [id, player]}
    <form on:submit|preventDefault={() => updatePlayer(id)}>
      <input type="string" bind:value={player.name} />
      <select bind:value={player.role}>
        {#each availableRoles as role}
          <option value={role}>{role}</option>
        {/each}
      </select>
      <select bind:value={player.node}>
        {#each Object.entries(graph.nodes) as [nodeID, _]}
          <option value={nodeID}>{nodeID}</option>
        {/each}
      </select>
      <button style="submit"> Update </button>
    </form>
  {/each}
  <Graph stateGraph={graph} />
  <p>Nodes</p>
  {#each Object.entries(graph.nodes) as [nodeID, node]}
    <div style="border: 1px solid black; margin: 2px">
      <b>{nodeID}</b>
      <button on:click={() => removeNode(nodeID)}> Remove </button>
      <div>
        {#if node.groups}
          Associated Groups:
          <ul>
            {#each node.groups as groupID}
              <li>
                {groupID}
                <button on:click={() => disassociateGroup(node, groupID)}>
                  Remove
                </button>
              </li>
            {/each}
          </ul>
        {/if}
        <form on:submit|preventDefault={() => associateGroup(nodeID)}>
          <select bind:value={groupsToAdd[nodeID]}>
            {#each Object.entries(graph.groups) as [groupID, _]}
              <option value={groupID}>{groupID}</option>
            {/each}
          </select>
          <button style="submit">Associate</button>
        </form>
      </div>
    </div>
  {/each}
  <button on:click={addNode}> addNode </button>
  <p>Groups</p>
  {#each Object.entries(graph.groups) as [groupID, group]}
    <div>
      {groupID}
      <button on:click={() => removeGroup(groupID)}> Remove </button>
    </div>
  {/each}
  <button on:click={addGroup}> addGroup </button>
  {#if facts.real}
    <p>Current Facts</p>
    {#each Object.entries(facts.real) as [name, fact]}
      <div>
        {name}: value:
        <select bind:value={fact.value}>
          {#each fact.possible as possible}
            <option value={possible}>{possible}</option>
          {/each}
        </select>
      </div>
    {/each}
    <button on:click={() => setRealFacts(facts.real)}> Update Values </button>
  {/if}
  Available Facts:
  {#each Object.entries(possibleFacts) as [name, fact]}
    <div>
      <label>
        <input type="checkbox" bind:checked={factsIncluded[name]} />
        {name}
      </label>
    </div>
  {/each}
  <button on:click={selectActiveFacts}>Select Facts </button>
  <button on:click={toggleConfig}>Back to game</button>
</main>
