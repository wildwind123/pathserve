<template>
  <nav class="navbar">
    <div >
        <div >
          <slot></slot>
        </div>
      </div>
  </nav>
</template>
<style scoped>
.navbar {
  display: flex;
  align-items: center;
  padding-left: 10px;
  background-color: rgb(136, 243, 243);
  height: 100%;
}
</style>