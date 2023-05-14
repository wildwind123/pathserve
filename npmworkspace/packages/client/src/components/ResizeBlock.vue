<template>
  <div >
    h:{{ height }} /  w:{{ width }}
  </div>
  <div :class="$style.resizable">
    <div ref="window" style="height: 100%; width: 100%; border: 1px solid black;">
      <slot></slot>
    </div>
    
  </div>
</template>
<script lang="ts" setup>

import { onMounted, ref, useCssModule } from "vue";
import interact from "interactjs"
import { useElementSize } from '@vueuse/core'

const window = ref<HTMLDivElement | null>(null)

const cssModule = useCssModule()
const {width, height} = useElementSize(window)

onMounted(() => {
  interact(`.${cssModule.resizable}`)
  .resizable({
    edges: { top: false, left: false, bottom: true, right: true },
    listeners: {
      move: function (event) {
        let { x, y } = event.target.dataset

        x = (parseFloat(x) || 0) + event.deltaRect.left
        y = (parseFloat(y) || 0) + event.deltaRect.top

        Object.assign(event.target.style, {
          width: `${event.rect.width}px`,
          height: `${event.rect.height}px`,
          transform: `translate(${x}px, ${y}px)`
        })

        Object.assign(event.target.dataset, { x, y })
      }
    }
  })
});
</script>
<style module>
.resizable {
  width: 500px;
  height: 500px;
  border-radius: 0.75rem;
  padding: 50px;
  background-color: rgb(255, 255, 255);
  box-sizing: border-box;
}
</style>
