<script>
 export let playerID;
 export let players;
 export let api;
 let role = "Train Depot";
 let uses = 1;
 console.log(uses)
 console.log
 function addKnownRole(role){
     var nodeID;
     for (const [_, player] of Object.entries(players)){
         if (player.role === role){
             nodeID = player.node;
         }
     }
     if (!nodeID){
         console.log("could not find a player with role ", role)
         return;
     }
     uses--;
     console.log(uses)
     let message = {
         event: "addKnownRole",
         playerID: playerID,
         nodeID: nodeID,
         role: role
     };
     api.socket.send(JSON.stringify(message));
 }
</script>
<main>
    {#if uses > 0}
        <button on:click={()=>addKnownRole(role)}>Reveal Train Expert</button>
    {:else}
        {#each Object.entries(players[playerID].knownRoles) as [nodeID, role]}
            <div>
                {nodeID} is {role}
            </div>
        {/each}
    {/if}
</main>
