import { defineStore } from 'pinia'

export const useWeatherStore = defineStore('weather', {
  state: () => ({file: ''}),
  getters: {
    dryBulbTemperature: () => 'not yet implemented',
  },
  actions: {
    read: () => 'not yet implemented',
  }
})