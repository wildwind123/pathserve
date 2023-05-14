import {defineConfig} from "vite"
import {resolve } from "pathe"
import dts from 'vite-plugin-dts'

export default defineConfig({
    plugins: [dts({
        insertTypesEntry: true,
    })],
    build: {
        lib: {
            entry: resolve(__dirname, 'src/Messenger.ts'),
            name: 'Messenger',
            fileName: 'messenger-lib',
            formats: ['es', 'cjs', 'umd', 'iife'],
        },
    }
})