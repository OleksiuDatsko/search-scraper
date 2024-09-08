<script lang="ts">
	import ListColumnItem from '$lib/components/ListItems/ListColumnItem.svelte';
	import { fetchLinksList } from '$lib/services/services';
	import { onMount } from 'svelte';
	import { links } from '$lib/utils/store';
	import type { Link } from '$lib/utils/types';

	export let list_type: string;
	let is_loading = false;

	onMount(async () => {
		is_loading = true;
		$links = [];
		fetchLinksList(list_type).then((data) => {
			is_loading = false;
			$links = data ? data : [];
		});
	});
</script>

<div class="table-container">
	<table class="table table-hover table-fixed">
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
			{#each $links as row (row.id)}
				<ListColumnItem link={row} listType={list_type} />
			{/each}
		</tbody>
	</table>
</div>
