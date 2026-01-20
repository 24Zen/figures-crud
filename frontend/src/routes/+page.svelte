<script lang="ts">
  import { onMount } from 'svelte';
  import type { Figure } from '$lib/types/Figure';
  import Gallery from '$lib/components/Gallery.svelte';

  let figures: Figure[] = [];
  let loading = true;
  let error = '';

  onMount(async () => {
    try {
      const res = await fetch('http://localhost:8080/figures');
      if (!res.ok) throw new Error('Failed to load figures');
      figures = await res.json();
    } catch (e) {
      error = e instanceof Error ? e.message : 'Unknown error';
    } finally {
      loading = false;
    }
  });
</script>

<h1 class="text-3xl font-bold text-center mb-8">
  Figure Collection
</h1>

{#if loading}
  <p class="text-center text-gray-500">Loading...</p>
{:else if error}
  <p class="text-center text-red-500">{error}</p>
{:else}
  <Gallery figures={figures} />
{/if}
