<script setup lang="ts">
import rpc from '@/rpc';
import { services } from '@/wailsjs/go/models'
import Collapsible from '@/components/Collapsible.vue'
</script>

<script lang="ts">

function getEPWHeader(epw: services.EPW): object {
  return {
    "City": epw.city,
    "Region": epw.region,
    "Country": epw.country,
    "Source": epw.source,
    "Station ID": epw.station_id,
    "Latitude": epw.latitude,
    "Longitude": epw.longitude,
    "Time Zone": epw.time_zone,
    "Elevation": epw.elevation,
  }

}

export default {
  data() {
    return {
      epwFile: '',
      epw: new services.EPW(),
      header: {},
    }
  },
  methods: {
    async openEPW() {
      this.epwFile = await rpc.app.OpenFileDialog("Select weather file")
      this.epw = await rpc.EPWService.Parse(this.epwFile)
      this.header = getEPWHeader(this.epw)
    },
  }
}

</script>

<template>
  <Collapsible prompt="open" text="Lorem ipsum" />
  <div class="content">
    <div class="epw">
      <div class="epw-header drawer" v-if="(epwFile != '')">
        <p>City: {{ epw.city }}</p>
        <p>Region, Country: {{ [epw.region, epw.country].join(", ") }}</p>
        <p>Latitude, longitude: {{ [epw.latitude, epw.longitude].join(", ") }}</p>
        <p>Elevation: {{ epw.elevation }}m</p>
      </div>
    </div>
    <div class="choose-epw">
      <p v-if="(epwFile == '')">no epw loaded</p>
      <button @click="openEPW">{{ (epwFile == '') ? "choose file" : "choose another file"}}</button>
    </div>
  </div>

</template>

<style lang="scss">
.content {
  position: absolute;
  padding: 10px;
  margin: 0;
  display: flex;
  flex-direction: column;
  width: 100%;
  align-content: flex-start;
}

.epw {
  flex-grow: 5;
}

.choose-epw {
  background-color: var(--accent3);
  width: 100%;
  left: 0px;
  bottom: 0px;
  border-radius: var(--border-medium);
  display: flex;
  align-items: center;
  padding: var(--border-medium);
  align-self: flex-end;

  p {
    padding: 2pt 10pt;
  }

  button {
    width: fit-content;
    height: fit-content;
    padding: 2pt 10pt;
    background: var(--light2);
    border-radius: var(--border-small);
    flex-grow: 0;
  }

  :last-child {
    margin-left: auto;
  }
}
</style>
