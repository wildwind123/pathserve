<template>
  <hr>
  <p >input 1</p>
  <button @click="setValue('button')">change input 1 to 'button'</button>
  <input
    type="text"
    :value="testValue"
    @input="setValue(($event.target as HTMLInputElement).value)"
  />
  <hr>
  <p>input 2</p>
  <input
    type="text"
    :value="testValue2"
    @input="s.setValue(($event.target as HTMLInputElement).value)"
  />
  <hr>
  <p>input 3</p>
  <input
    type="number"
    :value="testNumberValue"
    @input="setTestNumber(parseInt(($event.target as HTMLInputElement).value))"
  />
  <hr>
  <label>input 4</label>
  <pre>
        {{ testObject }}
    </pre
  >
  <button
    @click="
      () => {
        testObject = {
          ...testObject,
          id: testObject.id + 1,
          value: testObject.id + testObject.value,
        };
        setObject(testObject);
      }
    "
  >
    change testObject
  </button>
  <hr>
  <p for="">control button clicked</p>
  <pre>
    {{ buttonClickedCount }}
  </pre>
  <button @click="setButtonName('new button name')">change button name</button>
  <hr>
</template>
<script lang="ts" setup>
import {
  useString,
  useNumber,
  useObject,
  setListener,
  useButton,
} from "@pathserve/messenger";
import { onMounted, ref, toRaw } from "vue";

interface testObject {
  id: number;
  value: string;
}
// input 1
const testValue = ref("input 1");
const { setValue } = useString(testValue.value, "input 1", (value) => {
  testValue.value = value;
});
//  input 2
const testValue2 = ref(" input 2");
const s = useString(
  testValue2.value,
  "input 2",
  (value) => (testValue2.value = value)
);
// input 3
const testNumberValue = ref(0);
const { setValue: setTestNumber } = useNumber(
  testNumberValue.value,
  "input 3",
  (value) => (testNumberValue.value = value)
);
//  input 4
const testObject = ref<testObject>({ id: 1, value: "input 4" });
const { setValue: setObject } = useObject(
  toRaw(testObject.value),
  "input 4",
  (value) => (testObject.value = value as testObject)
);
// button 1
const buttonClickedCount = ref(0)
const {setButtonName} = useButton("button", () => {
    buttonClickedCount.value = buttonClickedCount.value + 1
})


onMounted(() => {
  setListener();
  window.addEventListener("message", (m) => { console.log(m) }, false);
});
</script>
