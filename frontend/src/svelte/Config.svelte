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
  export let showSecretDetails = true;
  let factsIncluded = {};
  let knownFacts = {};

  function shuffleArray(array) {
    for (let i = array.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [array[i], array[j]] = [array[j], array[i]];
    }
  }

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
  function toggleDetails() {
    showSecretDetails = !showSecretDetails;
  }
  function prepareSixPlayerGame() {
    let facts = {};
    for (const [name, fact] of Object.entries(possibleFacts)) {
      facts[name] = fact;
      factsIncluded[name] = true;
    }
    setRealFacts(facts);

    let nodes = Object.keys(graph.nodes);
    shuffleArray(nodes);

    let roles = []
    for (const [role, number] of Object.entries(availableRoles)) {
      for(let i=0; i < number; i++){
        roles.push(role);
      }
    }
    shuffleArray(roles);

    let factsInUse = Object.keys(factsIncluded);
    shuffleArray(factsInUse);

    for (const [id, player] of Object.entries(players)) {
      player.node = nodes.pop();
      player.role = roles.pop();
      updatePlayer(id);

      if (roleHasFact(player.role)) {
        knownFacts[id] = factsInUse.pop();
        setKnownFact(id, knownFacts[id]);
      } else {
        knownFacts[id] = "";
        clearKnownFact(id);
      }
    }
  }
  function roleHasFact(role) {
    const rolesWithFact = ["Fact Checker","Train Expert","Headquarters"]
    return rolesWithFact.includes(role);
  }
  function setRealFacts(facts) {
    let message = {
      event: "setRealFacts",
      facts: facts,
    };
    api.send(message);
  }
  function setKnownFact(id, factName) {
    let message = {
      event: "setKnownFact",
      playerID: id,
      names: [factName],
    };
    api.send(message);
  }
  function clearKnownFact(id) {
    let message = {
      event: "clearKnownFact",
      playerID: id,
    };
    api.send(message);
  }
  function updateRoles(){
      let message = {
          event: "updatePossibleRoles",
          possibleRoles: availableRoles,
      };
      api.send(message);
  }
</script>

<main>
  <p>Config</p>
  <p>Roles:</p>
  {#each Object.entries(availableRoles) as [role, count]}
      <div>
        {role}
        <input type="number" bind:value={availableRoles[role]} on:change={updateRoles} min=0>
      </div>
  {/each}
  <button on:click={() => toggleDetails()}>Toggle Details</button>
  <button on:click={() => prepareSixPlayerGame()}>dO aLl ThE tHiNgS</button>
  {#if showSecretDetails}
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
          {#each Object.entries(availableRoles) as [role, _]}
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
      <button on:click={() => setKnownFact(id, knownFacts[id])}>Update Known Fact
      </button>
    {/each}
    <Graph stateGraph={graph} {api} />
    <p>Nodes</p>
    {#each Object.entries(graph.nodes) as [nodeID, node]}
      <div style="border: 1px solid black; margin: 2px">
        <b>{nodeID}</b>
        <button on:click={() => removeNode(nodeID)}> Remove </button>
        <div>
          {#if Object.keys(node.groups).length > 0}
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
          {#if Object.keys(graph.groups).length > 0}
            <form on:submit|preventDefault={() => associateGroup(nodeID)}>
              <select bind:value={groupsToAdd[nodeID]}>
                {#each Object.entries(graph.groups) as [groupID, _]}
                  <option value={groupID}>{groupID}</option>
                {/each}
              </select>
              <button style="submit">Associate</button>
            </form>
          {/if}
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
  {/if}
  <button on:click={toggleConfig}>Back to game</button>
</main>
