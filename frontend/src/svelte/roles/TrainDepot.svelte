<script>
 export let selectedFacts;
 export let api;
 function guessFacts(){
     console.log("guessing")
     let message = {
         event: "submitFacts",
         submission: selectedFacts
     };
     api.send(message);
 }
</script>
<main>
    <button class="button is-danger" on:click={guessFacts}> guess the facts </button>
    <br/>
    CLICKING THIS BUTTON WILL END THE GAME!
</main>
