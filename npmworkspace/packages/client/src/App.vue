<template>
  <div class="app">
    <div class="app__header">
      <NavbarVue>
        <template v-if="app.selectedParam.value">
          <a :href="paramUrl(app.selectedParam.value)" target="blank"
            >Open in new tab.
            <strong>{{
              helper.getFileName(app.selectedParam.value.path)
            }}</strong></a
          >
        </template>
      </NavbarVue>
    </div>
    <div class="app__main">
      <aside class="app__menu">
        <MenuListVue
          @selected-key="app.selectedKey.value = $event"
          :items="app.info.params"
          :selected-key="app.selectedParam.value?.key ?? ''"
        >
        </MenuListVue>
      </aside>
      <div class="app__main-content-wrapper">
        <template v-if="app.selectedParam.value">
          <ResizeBlock>
            <iframe
              ref="iframe"
              width="100%"
              height="100%"
              :src="paramUrl(app.selectedParam.value)"
              frameborder="0"
            ></iframe>
          </ResizeBlock>
        </template>
      </div>
      <div class="app__main-control-wrapper"> 
        <Control :key="updateControl" @update:message="setSendMessage($event)" :messages="messages"></Control>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { useApp, paramUrl } from "@/composed/app";
import { onMounted, ref } from "vue";
import MenuListVue from "@/components/MenuList.vue";
import NavbarVue from "@/components/Navbar.vue";
import ResizeBlock from "@/components/ResizeBlock.vue";
import helper from "@/helper/helper";
import { Message } from "@pathserve/messenger";
import Control from "@/components/Control.vue";
import "@/style/bulma.css";
import "@/style/vars.css"
import _ from "lodash";
const app = useApp();
const iframe = ref<HTMLIFrameElement | null>(null);

const messages = ref<Message[]>([]);
const updateControl = ref(0)

onMounted(() => {
  window.addEventListener("message", receiveMessage, false);
  app.request();
});
function receiveMessage(event: MessageEvent<Message>) {
  if (!event.data.fromPathServe) {
    return;
  }
  const index = _.findIndex(messages.value, ["key", event.data.key]);
  if (index == -1) {
    messages.value.push(event.data);
    return;
  }
  // console.log('new value', event.data)
  messages.value[index] = event.data;
  updateControl.value = updateControl.value +1
}

function setSendMessage(message : Message) {
  const index = _.findIndex(messages.value, ["key", message.key]);
  if (index == -1) {
    console.warn(`message doesn't exist message = `, message, ',messages=' ,messages.value)
   return
  }
  messages.value[index] = message
  sendMessage(message)
}
function sendMessage(message : Message) {
      iframe.value!.contentWindow!.postMessage(_.cloneDeep(message), '*');
}


</script>
<style scoped>
.app {
  display: flex;
  flex-direction: column;
  height: 100vh;
  widows: 100vw;
}
.app__header {
  /* background-color: red; */
  height: var(--header-height);
  /* overflow: hidden; */
}
.app__main {
  /* flex: 1; */
  /* background-color: blue; */
  height: calc(100vh - var(--header-height));
  display: flex;
}
.app__menu {
  min-width: 200px;
  max-width: 220px;
  padding: 10px;
  /* border: 1px black solid ; */
  box-sizing: border-box;
  /* background-color: blueviolet; */
  height: calc(100vh - var(--header-height)) ;
  overflow: auto;
}
.app__main-content-wrapper {
  /* width: 500px;
  height: 500px; */
  padding: 20px;
  overflow: auto;
  background-color: aqua;
  flex: 1;
}
.app__main-control-wrapper {
  padding: 15px;
  height: calc(100vh - var(--header-height));
  overflow: auto;
}
</style>
