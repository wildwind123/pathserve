<template>
    test asdfasdfas
    <button @click="setValue('button')">button</button>

    <input type="text" :value="testValue" @input="setValue(($event.target as HTMLInputElement).value)">
    <input type="text" :value="testValue2" @input="s.setValue(($event.target as HTMLInputElement).value)">
</template>
<script lang="ts" setup>
import {useString, setListener} from "@pathserve/messenger"
import { onMounted, ref } from "vue";

const testValue = ref('asdfas')
const testValue2 = ref('asdfas2')

const {setValue, } = useString(testValue.value, "label 1", (value) => testValue.value = value)
const s = useString(testValue2.value, "label 2", (value) => testValue2.value = value)
window.pathServeMessageListener.callHooks.length
onMounted(() => {
    setListener()
})
</script>