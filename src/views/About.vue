<template>
  <div class="about">
    <Navbar class="navbar"/>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'

export default {
  name: "About",
  components: {
    Navbar
  }
}
</script>

<style scoped>
.about {
  padding: 0;
  margin: 0;
  position: absolute;
  left: 0;
  top: 0;
}

.navbar {
  position: sticky;
  top: 0;
  left: 0;
  padding: 0;
  margin: 0;
}

</style>
