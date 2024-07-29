<script lang="ts">
	import ListColumnItem from '$lib/components/ListItems/ListColumnItem.svelte';
import type { Link } from '$lib/utils/types';
	import { onMount } from 'svelte';

	export let list_type: string;
	let list: Link[] = [];
	let is_loading = false;

	async function fetchSearchResults(): Promise<Link[]> {
		is_loading = true;
		try {
			const response = await fetch(`http://localhost:8080/${list_type}`);
			if (response.ok) {
				const data: Link[] = await response.json();
				return data;
			} else {
				console.log('Error:', response.status);
				throw new Error('Failed to fetch search results');
			}
		} catch (error) {
			console.error(error);
			throw error;
		} finally {
			is_loading = false;
		}
	}

	onMount(() => {
		fetchSearchResults().then((data) => {
			list = data;
		});
	});
</script>

<div class="table-container">
	<table class="table table-hover table-fixed" >
		<thead>
			<tr>
				<th class="w-1/12">Id</th>
				<th class="w-3/12">Domain</th>
				<th>URL</th>
			</tr>
		</thead>
		<tbody>
			{#if is_loading}
				<p>Loading...</p>
			{/if}
			{#each list as row (row.id)}
				<ListColumnItem link={row} listType={list_type}/>
			{/each}
		</tbody>
	</table>
</div>
