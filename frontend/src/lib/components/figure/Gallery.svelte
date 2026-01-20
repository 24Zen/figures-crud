<script lang="ts">
  import type { Figure } from '$lib/types/Figure';
  import FigureInfo from './FigureInfo.svelte';

  export let figures: Figure[];
</script>

<div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
  {#each figures as figure}
    <FigureInfo {figure} />
  {/each}
</div>
