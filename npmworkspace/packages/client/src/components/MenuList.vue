<template>
  <div class="bulma-box">
    <aside class="bulma-menu">
      <template v-for="(item, key) in computedItems">
        <template v-if="key != 'from_config'">
          <p class="bulma-menu-label">{{ key }}</p>
          <ul class="bulma-menu-list">
            <template v-for="param in item">
              <template v-if="param.handler_config != ''">
                <li>
                <a
                  :class="{
                    ['bulma-is-active']: param.key == props.selectedKey,
                  }"
                  @click="emit('selected-key', param.key)"
                  >{{ param.fileName }}</a
                >
              </li>
              </template>
            </template>
          </ul>
        </template>
      </template>
    </aside>
  </div>
</template>
<script lang="ts" setup>
import { Param } from "@/composed/model";
import { computed } from "vue";
import _ from "lodash";
import helper from "@/helper/helper"

const props = defineProps<{
  items: Param[];
  selectedKey: string;
}>();
const emit = defineEmits<{
  (e: "selected-key", key: string): void;
}>();
interface Item {
  fileName: string;
  key: string;
  handler_config: string;
}
interface Items {
  [key: string]: Item[];
}

const computedItems = computed(() => {
  const items: Items = {};
  const sortedItems = _.sortBy(props.items, ["key"]);
  for (let i = 0; i < sortedItems.length; i++) {
    if (items[sortedItems[i].handler_config] === undefined) {
      items[sortedItems[i].handler_config] = [];
    }
    items[sortedItems[i].handler_config].push({
      fileName: helper.getFileName(sortedItems[i].path),
      key: sortedItems[i].key,
      handler_config: sortedItems[i].handler_config,
    });
  }
  return items;
});
</script>
<script></script>
