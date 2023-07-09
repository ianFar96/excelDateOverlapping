<script setup lang="ts">
import { ref } from 'vue';
import { Backend } from '../services/backend';
import { FileConflict } from '../../types/window';

const backend = new Backend();

const path = ref(localStorage.getItem('path') || '');
const date = ref(localStorage.getItem('date') || '');
const startTime = ref(localStorage.getItem('startTime') || '');
const endTime = ref(localStorage.getItem('endTime') || '');

const sheets = ref<FileConflict[]>([]);
const scan = async () => {
	try {
		if (!path.value) {
			alert('Path is empty');
			return;
		}

		localStorage.setItem('path', path.value);
		localStorage.setItem('date', date.value);
		localStorage.setItem('startTime', startTime.value);
		localStorage.setItem('endTime', endTime.value);

		sheets.value = await backend.scan(
			path.value,
			date.value,
			startTime.value,
			endTime.value
		);
	} catch (error) {
		alert(error);
	}
};
</script>

<template>
	<div class="bg-gray-900 min-h-screen min-w-screen w-full p-10">
		<form @submit.prevent="scan" class="bg-gray-800 text-white rounded-md p-6 block mb-10">
			<div class="mb-4">
				<label class="bold block mb-2">File path (excel)</label>
				<input type="text" class="w-full rounded-md px-3 py-2 outline-none bg-gray-700" v-model="path">
			</div>
		
			<div class="mb-4">
				<label class="bold block mb-2">Date column (ex. C)</label>
				<input type="text" class="w-full rounded-md px-3 py-2 outline-none bg-gray-700" v-model="date">
			</div>
		
			<div class="mb-4">
				<label class="bold block mb-2">Start time column (ex. D)</label>
				<input type="text" class="w-full rounded-md px-3 py-2 outline-none bg-gray-700" v-model="startTime">
			</div>
		
			<div class="mb-4">
				<label class="bold block mb-2">End time column (ex. E)</label>
				<input type="text" class="w-full rounded-md px-3 py-2 outline-none bg-gray-700" v-model="endTime">
			</div>
		
			<div class="flex justify-end">
				<button type="submit" class="rounded-md bg-gray-700 px-6 py-2">Scan</button>
			</div>
		</form>
	
		<div class="bg-gray-800 text-white rounded-md p-6 block" v-if="sheets.length > 0">
			<h1 class="text-3xl font-bold mb-4">
				Results
			</h1>
			<ul class="mb-4">
				<li v-for="sheet in sheets" :key="sheet.First">
					<h2 class="mb-4 mt-4 text-lg">Sheet name: {{ sheet.First }}</h2>
					<table class="w-full">
						<tr class="border border-white">
							<th v-for="col in sheet.Second" :key="col" class="whitespace-nowrap px-4 text-center">
								{{ col }}
							</th>
						</tr>
						<template v-for="conflict, index in sheet.Third" :key="index">
							<tr v-for="values, index in conflict" :key="index">
								<td class="px-4 text-center pt-4" v-for="value, index in values" :key="index">
									{{ value }}
								</td>
							</tr>
							<br>
						</template>
					</table>
					<hr>
				</li>
			</ul>
		</div>
	</div>
</template>