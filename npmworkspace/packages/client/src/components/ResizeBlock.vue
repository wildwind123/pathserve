<template>
  <div class="resize-block">
    <div class="resize-block__top-info">
        w:{{ resizeHelper.data.info.width }} h: {{ resizeHelper.data.info.height }}
    </div>
    <div ref="el" class="resize-block__bottom-rigth-resizer">WH</div>
    <div
      class="resize-block__block"
      :style="{
        width: `${resizeHelper.data.info.width}px`,
        height: `${resizeHelper.data.info.height}px`,
      }"
    >
      <slot></slot>
    </div>
  </div>
</template>
<script lang="ts" setup>
import useResizeHelper from "@/composed/resizeHelper";
import { onMounted, ref } from "vue";

const resizeHelper = useResizeHelper();

const el = ref<HTMLDivElement | null>(null);
onMounted(() => {
  resizeHelper.set(el.value!, {
    initialHeight: 500,
    initialWidth: 500,
  });
});
</script>
<style scoped>
.resize-block {
  display: inline-block;
  padding: 0 40px 40px 40px;
  background-color: rgb(255, 255, 255);
  border-radius: 50px;
  position: relative;
}
.resize-block__block {
  background-color: rgb(128, 136, 136);
  border: 1px rgb(192, 187, 187) solid;
}
.resize-block__bottom-rigth-resizer {
  position: absolute;
  bottom: 0;
  right: 0;
  background-color: beige;
  cursor: nw-resize;
}
.resize-block__top-info {
    height: 40px;
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>
