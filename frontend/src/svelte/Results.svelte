<script>
 export let facts;
 export let players;
 let right = '\u2714\uFE0F'
 let wrong = '\u274C'
</script>
<main>
    <div>
        Facts:
        <div>
            Submitted:
            {#each Object.entries(facts.real) as [name, {value}]}
                <div>
                Guessed: {facts.submitted[name]} Actual: {value} {facts.submitted[name] === value ? right : wrong}
                </div>
            {/each}
        </div>
    </div>
    <div>
        Roles:
        {#each Object.entries(players) as [_, {name, role}]}
            <div>
                {name} was {role}
            </div>
        {/each}
</main>
