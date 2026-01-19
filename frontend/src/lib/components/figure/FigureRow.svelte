<script lang="ts">
  import type { Figure } from '$lib/types/Figure';
  import FigureCard from './FigureCard.svelte';

  export let figures: Figure[];
</script>

<div class="flex gap-6 overflow-x-auto pb-4">
  {#each figures as figure}
    <FigureCard {figure} />
  {/each}
</div>
