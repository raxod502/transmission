<script>
 export let selectedFacts;
 export let api;
 function guessFacts(){
     console.log("guessing")
     let message = {
         event: "submitFacts",
         submission: selectedFacts
     };
     api.socket.send(JSON.stringify(message));
 }
</script>
<main>
    <button on:click={guessFacts}> guess the facts </button>
    CLICKING THIS BUTTON WILL END THE GAME!
</main>
