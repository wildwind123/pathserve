<template>
  <div class="app">
    <div class="app__header">
      <NavbarVue>
        <template v-if="app.selectedParam.value">
          <a :href="paramUrl(app.selectedParam.value)" target="blank">Open in new tab. <strong>{{ helper.getFileName(app.selectedParam.value.path) }}</strong></a>
        </template>
        
      </NavbarVue>
    </div>
    <div class="app__main">
      <aside class="app__menu">
        <MenuListVue @selected-key="app.selectedKey.value = $event" :items="app.info.params" :selected-key="app.selectedParam.value?.key ?? ''">

        </MenuListVue>
      </aside>
      <div class="app__main-content-wrapper">
        <template v-if="app.selectedParam.value">
          <iframe width="100%" height="90%" :src="paramUrl(app.selectedParam.value)"></iframe>
        </template>
        
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import {useApp, paramUrl} from "@/composed/app"
import { onMounted } from "vue";
import MenuListVue from "@/components/MenuList.vue"
import NavbarVue from "@/components/Navbar.vue";
import helper from "@/helper/helper";
import '@/style/bulma.css'
const app = useApp()

onMounted(() => {
  app.request()
})
</script>
<style scoped>
.app{
  display: flex;
  flex-direction: column;
  height: 100vh;
  widows: 100vw;
}
.app__header{
  /* background-color: red; */
}
.app__main {
  /* flex: 1; */
  /* background-color: blue; */
  height: 100%;
  display: flex;
}
.app__menu{
  min-width: 200px;
  max-width: 220px;
  margin: 10px;
  /* border: 1px black solid ; */
  box-sizing: border-box;
  /* background-color: blueviolet; */
}
.app__main-content-wrapper{
  /* width: 500px;
  height: 500px; */
  overflow: auto;
  background-color: aqua;
  flex: 1;
}
</style>
