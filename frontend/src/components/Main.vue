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
	<div>
		<label>File path (excel)</label>
		<input type="text" v-model="path">
	</div>

	<div>
		<label>Date column (ex. C)</label>
		<input type="text" v-model="date">
	</div>

	<div>
		<label>Start time column (ex. D)</label>
		<input type="text" v-model="startTime">
	</div>

	<div>
		<label>End time column (ex. E)</label>
		<input type="text" v-model="endTime">
	</div>

	<button type="button" @click="scan">Scan</button>

	<div v-if="sheets.length > 0">
		<ul>
			<li v-for="sheet in sheets" :key="sheet.First">
				<h2>{{ sheet.First }}</h2>
					<table>
						<tr>
							<th v-for="col in sheet.Second" :key="col">
								{{ col }}
							</th>
						</tr>
						<template v-for="conflict, index in sheet.Third" :key="index">
							<tr v-for="values, index in conflict" :key="index">
								<td v-for="value, index in values" :key="index">
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
</template>