<script lang="ts">
  import type { Figure } from '$lib/types/Figure';

  export let figure: Figure;
</script>

<div
  class="group relative bg-white rounded-2xl overflow-hidden
         shadow-md transition-all duration-300
         hover:shadow-2xl hover:-translate-y-1"
>
  <div class="overflow-hidden">
    <img
      src={figure.image}
      alt={figure.name}
      class="h-56 w-full object-cover
             transition-transform duration-300
             group-hover:scale-110"
    />
  </div>

  <div class="p-4">
    <h3 class="font-bold text-lg mb-1">
      {figure.name}
    </h3>

    <p class="text-sm text-gray-500 mb-2">
      {figure.anime}
    </p>

    <p class="text-sm text-gray-700 line-clamp-2">
      {figure.description}
    </p>
  </div>
</div>
