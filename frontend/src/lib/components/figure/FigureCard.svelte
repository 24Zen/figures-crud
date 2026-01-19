<script>
  import { goto } from '$app/navigation';
  export let figure;

  function goDetail() {
    goto(`/figures/${figure.id}`);
  }
</script>

<div
  class="w-60 bg-white rounded-lg shadow
         hover:shadow-lg hover:scale-[1.02]
         transition cursor-pointer"
  on:click={goDetail}
>
  <img
    src={figure.imageUrls[0]}
    alt={figure.character}
    class="h-56 w-full object-cover rounded-t-lg"
  />

  <div class="p-3">
    <h3 class="font-semibold">{figure.character}</h3>
    <p class="text-sm text-gray-600">{figure.anime}</p>
    <p class="text-xs text-gray-500 mt-1">{figure.description}</p>
  </div>
</div>
