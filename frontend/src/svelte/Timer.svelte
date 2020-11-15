<script>
  import { tweened } from "svelte/motion";
  export let endTime;
  export let startTime;
  export let api;
  let endDate = new Date(endTime);
  let startDate = new Date(startTime);
  let now = new Date();
  let originalLength = (endDate.getTime() - startDate.getTime()) / 1000;
  let timeLeft = (endDate.getTime() - now.getTime()) / 1000;

  let timer = tweened(timeLeft);
  let interval = setInterval(() => {
    if ($timer > 0) {
      $timer--;
    } else {
      gameOver();
    }
  }, 1000);

  $: minutes = Math.floor($timer / 60);
  $: minname = minutes > 1 ? "mins" : "min";
  $: seconds = Math.floor($timer - minutes * 60);
  function gameOver() {
    clearInterval(interval);
    let message = {
      event: "stopGame",
    };
    api.socket.send(JSON.stringify(message));
  }
</script>

<main>
  <h1>
    <span class="mins">{minutes}</span>:<span
      class="secs">{('' + seconds).padStart(2, '0')}</span>
    !
  </h1>
  <progress value={$timer / originalLength} />
</main>
