<script>
    import { tweened } from 'svelte/motion';
    export let endTime;
    export let startTime;
    let endDate = new Date(endTime);
    let startDate = new Date(startTime);
    let now = new Date();
    let originalLength = (endDate.getTime() - startDate.getTime())/1000;
 console.log(originalLength);
 console.log(now);
    let timeLeft = (endDate.getTime() - now.getTime())/1000;
 console.log(timeLeft);

    let timer = tweened(timeLeft)
  setInterval(() => {
    if ($timer > 0) $timer--;
  }, 1000);

  $: minutes = Math.floor($timer / 60);
  $: minname = minutes > 1 ? "mins" : "min";
  $: seconds = Math.floor($timer - minutes * 60)
</script>
<main>
<h1><span class="mins">{minutes}</span>:<span class="secs">{seconds}</span> !</h1>
<progress value={$timer/originalLength}></progress>
</main>
