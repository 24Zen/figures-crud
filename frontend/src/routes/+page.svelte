<script>
  import FigureInfo from '$lib/components/figure/FigureInfo.svelte';
  import Gallery from '$lib/components/figure/Gallery.svelte';

  let figures = [];

  fetch('http://localhost:8080/figures')
    .then(res => res.json())
    .then(data => figures = data);
</script>

{#each figures as f}
  <div class="border rounded mb-4 p-4">
    <FigureInfo
      title={f.anime}
      name={f.name}
      role={f.role}
    />
    <Gallery images={f.images} />
  </div>
{/each}
