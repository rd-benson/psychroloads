<script setup lang="ts">
import rpc from '@/rpc';
import { def } from '@vue/shared';




</script>

<script lang="ts">

export default {
  data() {
    return {
      epwFile: '',
      epw: rpc.epw
    }
  },
  methods: {
    async openFileDialog() {
      this.epwFile = await rpc.app.OpenFileDialog("Select weather file")
    }
  }
}

</script>

<template>
  <!-- Information -->
  <div class="content">
    <p @click="openFileDialog">choose file</p>
    <p v-if="(epwFile != '')">{{ epwFile.slice(0,50) }} </p>
  </div>
</template>

<style lang="scss">
.about {
  .title {
    margin: 30px auto 10px;
    font-size: 38px;
    color: #a150b5;
    text-align: center;
  }

  .content {
    position: relative;
    margin: 36px 20px;

    .comeon {
      position: absolute;
      left: 26px;
      top: 38px;
      max-width: 66%;

      img {
        width: 220px;
        height: 180px;
      }
    }

    .info {
      margin: 0 0 0 33%;
      font-size: 24px;
      text-align: left;

      .info-item {
        margin-bottom: 10px;

        .name {
          line-height: 40px;
          font-size: 28px;
          color: #6d6363;
        }

        .link {
          line-height: 30px;
          font-size: 20px;
          color: #5f6c86;
        }
      }
    }
  }

  .thank {
    height: 68px;
    line-height: 68px;
    margin: 36px auto;
    text-align: center;
    font-size: 40px;
  }
}
</style>
