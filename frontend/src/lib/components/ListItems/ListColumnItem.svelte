<script lang="ts">
	import type { Link } from '$lib/utils/types';
	import type { ModalSettings } from '@skeletonlabs/skeleton';

	import { getModalStore } from '@skeletonlabs/skeleton';

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
				link.url = r;
				updateLink();
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
				updateLink();
			}
		}
	};

	const modal_delete: ModalSettings = {
		type: 'confirm',
		title: 'Please Confirm',
		body: 'Are you sure you wish to delete this link?',
		response: (r: boolean) => {
			if (r) {
				deleteLink();
			}
		}
	};

	function updateLink() {
		fetch(`http://localhost:8080/${listType}/${link.id}`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ domain: link.domain, url: link.url, filterType: link.filter_type })
		});
	}

	function deleteLink() {
		fetch(`http://localhost:8080/${listType}/${link.id}`, {
			method: 'DELETE'
		});
	}
</script>

<tr>
	<td on:click={() => modalStore.trigger(modal_delete)}>{link.id}</td>
	<td on:click={() => modalStore.trigger(modal_domain)}>{link.domain}</td>
	<td class="truncate" on:click={() => modalStore.trigger(modal_url)}>{link.url}</td>
</tr>
