<template>
  <hr>
  <p >input 1</p>
  <p>
    <button @click="sendInput1('button')">change input 1 to 'button'</button>
  </p>
  <input
    type="text"
    :value="testValue"
    @input="sendInput1(($event.target as HTMLInputElement).value)"
  />
  <hr>
  <p>input 2</p>
  <input
    type="text"
    :value="testValue2"
    @input="sendInput2(($event.target as HTMLInputElement).value)"
  />
  <hr>
  <p>input 3</p>
  <input
    type="number"
    :value="testNumberValue"
    @input="sendInput3(parseInt(($event.target as HTMLInputElement).value))"
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
        sendInput4(testObject);
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
const [sendInput1] = useString("input 1",testValue.value, (value) => {
  testValue.value = value;
});
//  input 2
const testValue2 = ref(" input 2");
const [sendInput2] = useString(
  "input 2",
  testValue2.value,
  (value) => (testValue2.value = value)
);
// input 3
const testNumberValue = ref(0);
const [ sendInput3 ] = useNumber(
  "input 3",
  testNumberValue.value,
  (value) => (testNumberValue.value = value)
);
//  input 4
const testObject = ref<testObject>({ id: 1, value: "input 4" });
const [sendInput4] = useObject(
  "input 4",
  toRaw(testObject.value),
  (value) => (testObject.value = value as testObject)
);
// button 1
const buttonClickedCount = ref(0)
useButton("button", () => {
    buttonClickedCount.value = buttonClickedCount.value + 1
})


onMounted(() => {
  setListener();
});
</script>
