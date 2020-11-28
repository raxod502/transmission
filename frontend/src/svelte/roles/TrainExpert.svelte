<script>
 export let playerID;
 export let players;
 export let powerUses;
 export let nodes;
 export let api;
 export let role = "Train Depot";
 const uses = 1;
 console.log(uses)
 console.log(players[playerID].knownRoles)
 function addKnownRole(role){
     let nodeIDs = [];
     for (const [_, player] of Object.entries(players)){
         if (player.role === role && player.id != playerID){
             nodeIDs.push(player.node);
         }
     }
     if (nodeIDs.length == 0){
         console.log("could not find a player with role ", role)
         return;
     }
     for (const nodeID of nodeIDs){
        let message = {
            event: "addKnownRole",
            playerID: playerID,
            nodeID: nodeID,
            role: role
        };
        api.send(message);
     }
 }
</script>
<main>
    {#if uses > powerUses}
        <button class="button is-primary" on:click={()=>addKnownRole(role)}>Reveal {role}</button>
    {:else}
        {#if players[playerID].knownRoles}
            {#each Object.entries(players[playerID].knownRoles) as [nodeID, role]}
                <div>
                    {nodes[nodeID].name} is {role}
                </div>
            {/each}
        {:else }
            Loading...
        {/if}
    {/if}
</main>
