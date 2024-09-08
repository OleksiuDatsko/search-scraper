<script lang="ts">
	import SearchListItem from '$lib/components/ListItems/SearchListItem.svelte';
	import { searchedResults } from '$lib/utils/store';
	import { fetchSearchResults } from '$lib/services/services';

	let search_query = '';
	let is_submited = false;
	let is_submit_error = false;
	let is_loading = false;
	let depth = 5;

	async function handleSubmit() {
		if (search_query !== '') {
			is_submited = true;
			is_submit_error = false;
			try {
				is_loading = true;
				$searchedResults = await fetchSearchResults(search_query, depth);
			} catch (error) {
				is_submit_error = true;
			} finally {
				is_loading = false;
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
			<input
				type="search"
				placeholder="Search..."
				bind:value={search_query}
				on:keydown={(e) => e.key === 'Enter' && handleSubmit()}
			/>
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
		{#if $searchedResults.bot_detected}
			<aside class="alert variant-filled-warning m-4">
				<div class="alert-message">
					<h3 class="h3">BOT DETECTED</h3>
					<p>Scraper bot detected, need to wait a little bit.</p>
				</div>
			</aside>
		{/if}
		{#if $searchedResults.scraped_link.length > 0}
			<div class="card p-4">
				<header class="card-header">
					<h2 class="h2">Results ({$searchedResults.result_rating.toPrecision(3)}%)</h2>
					<span>{$searchedResults.query} n: {$searchedResults.scraped_link.length}</span>
				</header>
				<section class="p-4">
					<ul class="list">
						{#each $searchedResults.scraped_link as link (link.id)}
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
