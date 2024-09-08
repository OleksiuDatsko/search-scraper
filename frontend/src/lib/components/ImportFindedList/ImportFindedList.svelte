<script lang="ts">
	import { FileButton, SlideToggle } from '@skeletonlabs/skeleton';

	let files: FileList;
    let overite = false;

	async function handleSubmit() {
		const formData = new FormData();
		formData.append('file', files[0]);
		const response = await fetch(`http://localhost:8080/findedlist/import`, {
			method: 'POST',
			body: formData
		});
        if (response.ok) {
            window.location.reload()
        }
	}
</script>

<div class="py-4 m-auto flex items-center justify-center gap-10">
	<FileButton
		name="files"
		button="btn variant-soft-primary"
		bind:files
		on:change={() => handleSubmit()}>Upload</FileButton
	>
    <SlideToggle name="slide" bind:checked={overite} >Overite (dosen't work)</SlideToggle>
</div>
