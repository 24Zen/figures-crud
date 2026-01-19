<script>
  export let images = [];
</script>

<div class="flex gap-2 overflow-x-auto">
  {#each images as img}
    <img src={img} alt="" class="h-32 rounded" />
  {/each}
</div>
