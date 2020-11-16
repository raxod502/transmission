<script>
 export let facts;
 export let playerID;
 export let api;
 function revealFacts(){
     let names = [];
     for (const [name, _] of Object.entries(facts.real)){
         names.push(name);
    }
     let message = {
         event: "addKnownFact",
         playerID: playerID,
         names: names
     }
     api.socket.send(JSON.stringify(message));
 }
</script>
<main>
    <button on:click={revealFacts}>Reveal Facts</button>
</main>
