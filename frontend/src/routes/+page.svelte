<script lang="ts">
	import SearchListItem from '$lib/components/ListItems/SearchListItem.svelte';
	import type { SearchResult } from '$lib/utils/types';

	let search_query = '';
	let search_results: SearchResult | null = null;
	let is_submited = false;
	let is_submit_error = false;
	let is_loading = false;
	let is_bot_detected = false;
	let depth = 5;

	$: console.log(search_query);

	async function fetchSearchResults(): Promise<SearchResult> {
		is_loading = true;
		try {
			const response = await fetch(`http://localhost:8080/search?q=${search_query}&d=${depth}`);
			if (response.ok) {
				const data: SearchResult = await response.json();
				is_bot_detected = response.status === 226;
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

	async function handleSubmit() {
		if (search_query !== '') {
			is_submited = true;
			is_submit_error = false;
			try {
				search_results = await fetchSearchResults();
			} catch (error) {
				is_submit_error = true;
				search_results = null;
			}
		} else {
			is_submit_error = true;
			is_submited = false;
		}
	}
</script>

<section>
	<div class="container m-auto">
		<h1 class="h1 text-center">Search</h1>
		<div class="mb-4 input-group input-group-divider grid-cols-[auto_1fr_auto_auto] bg-none">
			<div class="input-group-shim">S</div>
			<input type="search" placeholder="Search..." bind:value={search_query} />
			<input type="number" bind:value={depth} />
			<button
				class="variant-filled-secondary"
				class:input-success={is_submited}
				class:input-error={is_submit_error}
				on:click={handleSubmit}>Submit</button
			>
		</div>
	</div>
</section>

<section>
	<div class="container m-auto">
		{#if is_loading}
			<div class="alert variant-outline">Loading...</div>
		{/if}
		{#if search_results}
			{#if is_bot_detected}
				<aside class="alert variant-filled-warning m-4">
					<div class="alert-message">
						<h3 class="h3">BOT DETECTED</h3>
						<p>Scraper bot detected, need to wait a little bit.</p>
					</div>
				</aside>
			{/if}
			<div class="card p-4">
				<header class="card-header">
					<h2 class="h2">Results ({search_results.result_rating}%)</h2>
				</header>
				<section class="p-4">
					<ul class="list">
						{#each search_results.scraped_link as link (link.id)}
							<SearchListItem {link} />
						{/each}
					</ul>
				</section>
			</div>
		{/if}
		{#if is_submit_error}
			<div class="alert variant-filled-warning m-4">
				<div class="alert-message">
					<h3 class="h3">ERROR</h3>
					<p>Error fetching results. Please try again.</p>
				</div>
			</div>
		{/if}
	</div>
</section>
