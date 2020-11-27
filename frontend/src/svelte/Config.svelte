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
  let knownFacts = {};

  function updatePlayer(id) {
    let player = players[id];
    console.log("update player:", player);
    let playerMessage = {
      event: "updatePlayer",
      player: player,
    };
    api.socket.send(JSON.stringify(playerMessage));
    if (player.node) {
      let nodeMessage = {
        event: "updateNode",
        node: {
          ...graph.nodes[player.node],
          player: player.id,
          name: player.name,
          color: player.color,
        },
      };
      api.socket.send(JSON.stringify(nodeMessage));
    }
  }
  function addNode() {
    let message = {
      event: "updateNode",
      node: {
        id: "node-" + uuidv4(),
        groups: [],
      },
    };
    api.send(message);
  }
  function removeNode(id) {
    let message = {
      event: "removeNode",
      nodeID: id,
    };
    api.send(message);
  }
  function addGroup() {
    let message = {
      event: "updateGroup",
      group: {
        id: "group-" + uuidv4(),
        messages: [],
      },
    };
    api.send(message);
  }
  function removeGroup(id) {
    let message = {
      event: "removeGroup",
      groupID: id,
    };
    api.send(message);
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
    api.send(message);
  }
  function disassociateGroup(node, groupID) {
    let newGroups = node.groups.filter((g) => g !== groupID);
    let message = {
      event: "updateNode",
      node: {
        ...node,
        groups: newGroups,
      },
    };
    api.send(message);
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
    api.send(message);
  }
  function addKnownFact(id, factName) {
    let message = {
      event: "addKnownFact",
      playerID: id,
      names: [factName],
    };
    api.send(message);
  }
</script>

<main>
  <p>Config</p>
  <p>Players:</p>
  {#each Object.entries(players) as [id, player]}
    <form>
      <input
        type="string"
        bind:value={player.name}
        on:input={() => updatePlayer(id)} />
      <!-- svelte-ignore a11y-no-onchange -->
      <select bind:value={player.role} on:change={() => updatePlayer(id)}>
        <option value="">(no role)</option>
        {#each availableRoles as role}
          <option value={role}>{role}</option>
        {/each}
      </select>
      <!-- svelte-ignore a11y-no-onchange -->
      <select bind:value={player.node} on:change={() => updatePlayer(id)}>
        <option value="">(no node)</option>
        {#each Object.entries(graph.nodes) as [nodeID, _]}
          <option value={nodeID}>{nodeID}</option>
        {/each}
      </select>
    </form>
    <select bind:value={knownFacts[id]}>
      <option value="">(no fact)</option>
      {#each Object.entries(facts.real) as [name, _]}
        <option value={name}>{name}</option>
      {/each}
    </select>
    <button on:click={() => addKnownFact(id, knownFacts[id])}>Update Known Fact
    </button>
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
  {#each Object.keys(graph.groups) as groupID}
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
        {name}:

        <!-- svelte-ignore a11y-no-onchange -->
        <select
          bind:value={fact.value}
          on:change={() => setRealFacts(facts.real)}>
          {#each fact.possible as possible}
            <option value={possible}>{possible}</option>
          {/each}
        </select>
      </div>
    {/each}
  {/if}
  Available Facts:
  {#each Object.keys(possibleFacts) as name}
    <div>
      <label>
        <input
          type="checkbox"
          bind:checked={factsIncluded[name]}
          on:change={selectActiveFacts} />
        {name}
      </label>
    </div>
  {/each}
  <button on:click={toggleConfig}>Back to game</button>
</main>
