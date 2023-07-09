<script setup lang="ts">
import { Ref, ref } from 'vue';
import { Backend } from '../services/backend';
const backend = new Backend();

let sheets: Ref<string[]>;
try {
	sheets = ref(await backend.scan('path/to/file.xls','E1', 'E2', 'E3'));
} catch (error) {
	alert(error);
}
</script>

<template>
  {{sheets}}
</template>