<script>
 export let facts;
 export let playerID;
 export let api;
 export let checks;
 let selectedFact;
 let selectedValue;
 let uses = 1;
 function submitCheck(){
     let message = {
         event: "checkFact",
         playerID: playerID,
         field: selectedFact,
         value: selectedValue
     };
     api.socket.send(JSON.stringify(message));
     uses--;
 }
</script>
<main>
    {#if uses > 0}
    <div>
        Fact:
        <select bind:value={selectedFact}>
            {#each Object.entries(facts.real) as [name, _]}
                <option value={name}>
                    {name}
                </option>
            {/each}
        </select>
    </div>
    <div>
        {#if facts.real[selectedFact]}
            Value to check:
            <select bind:value={selectedValue}>
                {#each facts.real[selectedFact].possible as value}
                    <option value={value}>
                        {value}
                    </option>
                {/each}
            </select>
        {/if}
    </div>
    <button on:click={submitCheck} disabled={uses <= 0}> Check Fact </button>
    {:else}
            Results:
            {#each checks as {name, guessedValue, correct}}
                <div>
                    Guess {guessedValue} for {name} was {correct ? "right!" : "wrong"}
                </div>
            {/each}
    {/if}
</main>
