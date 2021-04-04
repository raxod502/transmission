<script>
  export let graph;
  export let players;
  export let playerID;
  export let getInterlocutorName;
  export let checkForSend;
  export let composedMessages;
  export let enabled = true;
</script>

<main>
    <div class="row" style="height: 80%">
        <div class="columns is-gapless mx-2" style="height: 100%">
            {#if Object.keys(graph.groups).length === 0}
            <p>No conversations</p>
            {:else if players[playerID].node === ""}
            <p>You weren't assigned to a node :(</p>
            {:else}
            {#each graph.nodes[players[playerID].node].groups as groupID}
                {#if graph.groups[groupID]}
                <div class="column" style="height: 100%">
                    <div class="rows" style="height: 100%; display: flex">
                    <div class="row mx-3" style="width: 100%; height: 100%">
                        <div
                        style="overflow-y: auto; height: calc(100% - 60px); display: flex">
                        <div style="margin-top: auto">
                            {#each graph.groups[groupID].messages.slice(-20) as { sender, text }}
                            <p
                                style="color: {graph.nodes[sender].color}"
                                class="mb-3">
                                {graph.nodes[sender].name}:
                                {text}
                            </p>
                            {/each}
                        </div>
                        </div>
                        <div style="height: 60px">
                        <input
                            type="text"
                            class="input"
                            disabled={!enabled}
                            on:keyup|preventDefault={() => checkForSend(groupID)}
                            bind:value={composedMessages[groupID]}
                            style="width: 100%"
                            placeholder="Message to {getInterlocutorName(playerID, groupID)}" />
                        </div>
                    </div>
                    </div>
                </div>
                {/if}
            {/each}
            {/if}
        </div>
    </div>
</main>
