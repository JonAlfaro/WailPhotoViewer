<script>
	import { getContext } from 'svelte';
	import ModalContent from './ModalContent.svelte'
	let name = 'Jon';

	
	const { open } = getContext('simple-modal');

	const handleOpenModal = (val) => {
		window.backend.helloSir(val).then((result) => {
			open(ModalContent, { message: result });
		});
	};
</script>

<main>
	<p>Input your name: <input bind:value={name}></p>
	<p><button on:click={handleOpenModal(name)}>Say Hello</button></p>
</main>

<style></style>