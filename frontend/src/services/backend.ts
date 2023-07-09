import { ExtendedWindow } from '../../types/window';

export class Backend {
	private window: ExtendedWindow;

	constructor() {
		this.window = window as ExtendedWindow;
	}

	scan(filePath: string, dateCol: string, startTimeCol: string, endTimeCol: string) {
		return this.window.scan(filePath, dateCol, startTimeCol, endTimeCol);
	}
}