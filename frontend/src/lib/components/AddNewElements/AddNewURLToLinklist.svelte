<script lang="ts">
	export let list_type: string;
	function handleSubmit() {
		fetch(`http://localhost:8080/${list_type}`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ url })
		});
	}

	let url = '';
</script>

<div class="m-4">
	<h3 class="h3 mx-6">Add new site (input URL)</h3>
	<div class="input-group input-group-divider grid-cols-[1fr_auto]">
		<input type="text" bind:value={url} placeholder="https://www.example.com" />
		<button class="variant-filled-secondary" on:click={handleSubmit}>Submit</button>
	</div>
</div>
