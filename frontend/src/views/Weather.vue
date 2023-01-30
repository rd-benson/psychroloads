<script setup lang="ts">
import rpc from '@/rpc';
import { services } from '@/wailsjs/go/models'
</script>

<script lang="ts">

export default {
  data() {
    return {
      epwFile: '',
      epw: new(services.EPW),
      header: '',
    }
  },
  methods: {
    async openFileDialog() {
      this.epwFile = await rpc.app.OpenFileDialog("Select weather file")
      this.epw = await rpc.EPWService.Parse(this.epwFile)
      this.header = await rpc.EPWService.Header(this.epw)
    }
  }
}

</script>

<template>
  <!-- Information -->
  <div class="content">
    <div>
      <p v-if="(epwFile == '')">no epw loaded</p>
      <button @click="openFileDialog">choose file</button>
    </div>
    <p v-if="(epwFile != '')">{{ header }}</p>
  </div>
</template>

<style lang="scss">
.content {
  display: flex;
  flex-direction: column;
  p, button {
    width: fit-content;
    padding: 2pt 10pt;
  }
  button {
    background: var(--light2);
    border-radius: 4pt;
  }
 
}
</style>
