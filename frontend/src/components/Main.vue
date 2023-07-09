<script setup lang="ts">
import { ref } from 'vue';
import { Backend } from '../services/backend';
const backend = new Backend();

const path = ref(localStorage.getItem('path') || '');

const sheets = ref<string[]>([]);
const scan = async () => {
	try {
		if (!path.value) {
			alert('Path is empty');
			return;
		}

		localStorage.setItem('path', path.value);

		sheets.value = await backend.scan(path.value, 'A', 'C', 'D');
	} catch (error) {
		alert(error);
	}
};
</script>

<template>
  {{sheets}}

	<input type="text" v-model="path">
	<button type="button" @click="scan">Scan</button>
</template>