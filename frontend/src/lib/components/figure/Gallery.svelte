<script>
  import FigureInfo from './FigureInfo.svelte';
  export let figures = [];
</script>

<div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
  {#each figures as figure}
    <FigureInfo {figure} />
  {/each}
</div>
