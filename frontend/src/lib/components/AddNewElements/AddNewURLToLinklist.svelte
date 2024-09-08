<script lang="ts">
	import { addLink } from '$lib/services/services';
	import { links } from '$lib/utils/store';

	export let list_type: string;
	let url = '';

	async function handleSubmit() {
		const link = await addLink(url, list_type);
		$links.push(link);
		$links = $links;
		console.log($links);
	}
</script>

<div class="m-4">
	<h3 class="h3 mx-6">Add new site (input URL)</h3>
	<div class="input-group input-group-divider grid-cols-[1fr_auto]">
		<input
			type="text"
			bind:value={url}
			on:keydown={(e) => e.key === 'Enter' && handleSubmit()}
			placeholder="https://www.example.com"
		/>
		<button class="variant-filled-secondary" on:click={handleSubmit}>Submit</button>
	</div>
</div>
