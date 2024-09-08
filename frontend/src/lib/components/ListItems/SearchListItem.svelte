<script lang="ts">
	import { blockedIds, findedIds } from '$lib/utils/store';
	import type { ScrapedLink } from '$lib/utils/types';

	let blockedIdsValue: number[];
	blockedIds.subscribe((v) => (blockedIdsValue = v));

	let findedIdsValue: number[];
	findedIds.subscribe((v) => (findedIdsValue = v));

	export let link: ScrapedLink;

	async function handleAddToList(list_type: string) {
		fetch(`http://localhost:8080/${list_type}`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				url: link.link,
				domain: link.domain
			})
		})
			.then((res) => res.json())
			.then((data) => {
				if (data.error) {
					console.error(data.error);
				}
			})
			.catch((error) => {
				console.error('Error:', error);
			});
	}
</script>

<li class="mb-4 flex flex-row items-center justify-between">
	<a href={link.link} target="_blank" class="flex flex-col w-3/5">
		<small>{link.domain}</small>
		<h3>{link.title}</h3>
		<p>{link.snipped}</p>
	</a>
	<div class="btn-group variant-filled">
		<button
			on:click={() => {
				blockedIds.update((ids) => [...ids, link.id]);
				handleAddToList('blacklist');
			}}
			disabled={blockedIdsValue.includes(link.id) || findedIdsValue.includes(link.id)}
			class:variant-filled-secondary={blockedIdsValue.includes(link.id)}>Block</button
		>
		<button
			on:click={() => {
				findedIds.update((ids) => [...ids, link.id]);
				handleAddToList('findedlist');
			}}
			disabled={blockedIdsValue.includes(link.id) || findedIdsValue.includes(link.id)}
			class:variant-filled-secondary={findedIdsValue.includes(link.id)}>Finded</button
		>
	</div>
</li>
<hr class="!border-t-2" />
