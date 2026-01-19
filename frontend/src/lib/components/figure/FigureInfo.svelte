<script>
  export let figure;
</script>

<div
  class="bg-white rounded-2xl overflow-hidden
         shadow-md cursor-pointer
         transition-all duration-300 ease-out
         hover:shadow-2xl
         hover:-translate-y-2
         hover:scale-[1.02]"
>
  <!-- รูป -->
  <div class="overflow-hidden">
    <img
      src={figure.imageUrls?.[0]}
      alt={figure.character}
      class="h-48 w-full object-cover
             transition-transform duration-300
             hover:scale-110"
    />
  </div>

  <!-- ข้อมูล -->
  <div class="p-4">
    <h2 class="font-bold text-lg mb-1">
      {figure.character}
    </h2>

    <p class="text-sm text-gray-500 mb-2">
      {figure.anime}
    </p>

    <p class="text-sm text-gray-700 line-clamp-3">
      {figure.description}
    </p>
  </div>
</div>
