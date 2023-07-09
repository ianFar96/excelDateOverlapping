export type ExtendedWindow = Window & typeof globalThis & {
  scan: (filePath: string, dateCol: string, startTimeCol: string, endTimeCol: string) => Promise<FileConflict[]>
}

export type FileConflict = {
  First: string
  Second: string[]
  Third: Conflict[]
}

export type Conflict = {
  First: string[],
  Second: string[]
}