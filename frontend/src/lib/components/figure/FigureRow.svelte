<script>
  import FigureCard from './FigureCard.svelte';
  export let figures = [];
</script>

<div class="flex gap-6 overflow-x-auto pb-4">
  {#each figures as figure}
    <FigureCard {figure} />
  {/each}
</div>
