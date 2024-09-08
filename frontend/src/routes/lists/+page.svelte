<script lang="ts">
	import { Tab, TabGroup } from '@skeletonlabs/skeleton';
	import ListListItem from '$lib/components/Tables/ListTable/ListTable.svelte';
	import ImportFindedList from '$lib/components/ImportFindedList/ImportFindedList.svelte';
	import AddNewUrlToWhitelist from '$lib/components/AddNewElements/AddNewURLToLinklist.svelte';
	import { tabSet } from '$lib/utils/store';
</script>

<section class="container m-auto card p-4">
	<TabGroup justify="justify-center">
		<Tab bind:group={$tabSet} name="findeded" value={0}>Finded list</Tab>
		<Tab bind:group={$tabSet} name="blacklist" value={1}>Blacklist</Tab>
		<Tab bind:group={$tabSet} name="whitelist" value={2}>Whitelist</Tab>
		<!-- Tab Panels --->
		<svelte:fragment slot="panel">
			{#if $tabSet === 0}
                <ImportFindedList />
				<ListListItem list_type="findedlist" />
			{:else if $tabSet === 1}
				<ListListItem list_type="blacklist" />
			{:else if $tabSet === 2}
				<AddNewUrlToWhitelist list_type="whitelist"/>
				<ListListItem list_type="whitelist" />
			{/if}
		</svelte:fragment>
	</TabGroup>
</section>
