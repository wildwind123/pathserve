<template>
    test asdfasdfas
    <button @click="setValue('button')">button</button>

    <input type="text" :value="testValue" @input="setValue(($event.target as HTMLInputElement).value)">
    <input type="text" :value="testValue2" @input="s.setValue(($event.target as HTMLInputElement).value)">
    <input type="number" :value="testNumberValue" @input="setTestNumber(parseInt(($event.target as HTMLInputElement).value))">

    <label>testObject</label>
    <pre>
        {{ testObject }}
    </pre>
    <button @click="() => {
        testObject = {...testObject, id : testObject.id+1 , value: testObject.id + testObject.value}
        setObject(testObject)
    }">change testObject</button>
</template>
<script lang="ts" setup>
import {useString, useNumber , useObject  , setListener} from "@pathserve/messenger"
import { onMounted, ref, toRaw } from "vue";

const testValue = ref('asdfas')
const testValue2 = ref('asdfas2')
const testNumberValue = ref(0)
interface testObject {
    id : number,
    value: string,
}
const testObject = ref<testObject>({id: 1, value: 'ss'})


const {setValue, } = useString(testValue.value, "label 1", (value) => {
    testValue.value = value
})
const s = useString(testValue2.value, "label 2", (value) => testValue2.value = value)
const { setValue : setTestNumber } =  useNumber(testNumberValue.value, "label 3", (value) => testNumberValue.value = value)

const {setValue : setObject } = useObject(toRaw(testObject.value), "label 4", (value) => testObject.value = value as testObject)


onMounted(() => {
    setListener()
})
</script>