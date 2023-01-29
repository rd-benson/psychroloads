import { useWeatherStore } from "../stores/weather";

export default {
    setup() {
        const store = useWeatherStore()

        return {
            store
        }
    }
}