<script lang="ts">
  import type { Figure } from '$lib/types/Figure';

  export let figure: Figure;
</script>

<div
  class="bg-white rounded-xl shadow
         hover:shadow-xl transition-all duration-300"
>
  <img
    src={figure.image}
    alt={figure.name}
    class="h-48 w-full object-cover rounded-t-xl"
  />

  <div class="p-4">
    <h2 class="font-semibold text-lg">
      {figure.name}
    </h2>

    <p class="text-sm text-gray-500 mb-2">
      {figure.anime}
    </p>

    <p class="text-sm text-gray-700">
      {figure.description}
    </p>
  </div>
</div>
