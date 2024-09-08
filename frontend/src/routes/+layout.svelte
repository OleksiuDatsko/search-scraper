<script lang="ts">
	import '../app.postcss';
	import { LightSwitch } from '@skeletonlabs/skeleton';
	import { initializeStores, Modal } from '@skeletonlabs/skeleton';

	initializeStores();
</script>

<header class="bg-surface-100-800-token p-4 flex items-center mb-4">
	<div class="container flex justify-between m-auto">
		<nav>
			<a href="/" class="btn btn-sm variant-filled">Search</a>
		</nav>
		<nav class="flex gap-2 items-center">
			<a href="/lists" class="btn btn-sm variant-filled">Lists</a>

			<LightSwitch />
		</nav>
	</div>
</header>
<Modal />
<slot />

