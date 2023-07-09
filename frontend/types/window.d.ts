export type ExtendedWindow = Window & typeof globalThis & {
  scan: (filePath: string, dateCol: string, startTimeCol: string, endTimeCol: string) => Promise<string>
}