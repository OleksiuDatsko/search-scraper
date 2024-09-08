<script lang="ts">
	import type { Link } from '$lib/utils/types';
	import type { ModalSettings } from '@skeletonlabs/skeleton';
	import { updateLink, deleteLink } from '$lib/services/services';
	import { getModalStore } from '@skeletonlabs/skeleton';
	import { SvelteComponent } from 'svelte';
	import { links } from '$lib/utils/store';

	const modalStore = getModalStore();

	export let link: Link;
	export let listType: string;

	const modal_domain: ModalSettings = {
		type: 'prompt',
		title: 'Enter Domain',
		body: 'Provide new domain in the field below.',
		value: link.domain,
		valueAttr: { type: 'text', minlength: 3, required: true },
		response: (r: string) => {
			if (r) {
				link.domain = r;
				updateLink(link, listType);
			}
		}
	};

	const modal_url: ModalSettings = {
		type: 'prompt',
		title: 'Enter URL',
		body: 'Provide new url in the field below.',
		value: link.url,
		valueAttr: { type: 'text', minlength: 3, required: true },
		response: (r: string) => {
			if (r) {
				link.url = r;
				updateLink(link, listType);
			}
		}
	};

	const modal_delete: ModalSettings = {
		type: 'confirm',
		title: 'Please Confirm',
		body: 'Are you sure you wish to delete this link?',
		response: (r: boolean) => {
			if (r) {
				deleteLink(link, listType);
				$links = $links.filter((l) => l.id !== link.id);
			}
		}
	};


</script>

<tr>
	<td on:click={() => modalStore.trigger(modal_delete)}>{link.id}</td>
	<td on:click={() => modalStore.trigger(modal_domain)}>{link.domain}</td>
	<td class="truncate" on:click={() => modalStore.trigger(modal_url)}>{link.url}</td>
</tr>
