<script>
	import { onMount } from "svelte";

	let figures = [];
	let loading = true;
	let error = "";

	onMount(async () => {
		try {
			const res = await fetch("http://localhost:8080/figures");
			if (!res.ok) {
				throw new Error("Failed to load figures");
			}
			figures = await res.json();
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	});
</script>

<h1 class="text-center text-4xl font-bold mb-8">
	Figure Collection
</h1>


{#if loading}
	<p>Loading...</p>
{:else if error}
	<p style="color: red">{error}</p>
{:else}
	<div class="grid">
		{#each figures as f}
			<div class="card">

				<!-- รูปหลัก -->
				<img
					class="main-img"
					src={f.imageUrls && f.imageUrls.length > 0
						? f.imageUrls[0]
						: "https://via.placeholder.com/200x250?text=No+Image"}
					alt={f.character}
				/>

				<!-- รูปย่อย 3 รูป -->
				{#if f.imageUrls && f.imageUrls.length > 1}
					<div class="thumbs">
						{#each f.imageUrls.slice(0, 3) as img}
							<img src={img} alt="thumb" />
						{/each}
					</div>
				{/if}

				<h3>{f.character}</h3>
				<p>{f.anime}</p>
				<small>{f.description}</small>
			</div>
		{/each}
	</div>
{/if}

<style>
	h1 {
		margin-bottom: 20px;
	}

.grid {
	display: grid;
	grid-template-columns: repeat(3, 1fr);
	gap: 24px;
	max-width: 1100px;
	margin: 0 auto;
}

@media (max-width: 1024px) {
	.grid {
		grid-template-columns: repeat(2, 1fr);
	}
}

@media (max-width: 640px) {
	.grid {
		grid-template-columns: 1fr;
	}
}


	.card {
		background: white;
		padding: 10px;
		border-radius: 8px;
		box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
		overflow: hidden;
		transition: transform 0.2s ease, box-shadow 0.2s ease;
		cursor: pointer;
        display: flex;
	    flex-direction: column;
        }

    small {
	    min-height: 32px;
    }

	.card:hover {
		transform: translateY(-6px);
		box-shadow: 0 8px 20px rgba(0, 0, 0, 0.18);
	}

	/* รูปหลัก */
	.main-img {
		width: 100%;
		height: 230px;
		object-fit: cover;
		border-radius: 6px;
		margin-bottom: 6px;
		transition: transform 0.3s ease;
	}

	.card:hover .main-img {
		transform: scale(1.05);
	}

	/* รูปย่อย */
	.thumbs {
		display: flex;
		gap: 6px;
		margin-bottom: 6px;
	}

	.thumbs img {
		flex: 1;
		height: 60px;
		object-fit: cover;
		border-radius: 4px;
		opacity: 0.85;
		transition: transform 0.2s ease, opacity 0.2s ease;
	}

	.thumbs img:hover {
		transform: scale(1.05);
		opacity: 1;
	}

	h3 {
		margin: 8px 0 4px;
	}

	.card:hover h3 {
		color: #2563eb;
	}

	p {
		margin: 0;
		color: #555;
	}

	small {
		color: #777;
	}
</style>
