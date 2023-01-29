import { createApp } from "vue";
import { createPinia } from "pinia";
import piniaPersistedState from 'pinia-plugin-persistedstate'

import router from "@/router/router";

import App from "@/App.vue"

import "./style.scss";

// import "./assets/main.css";

const pinia = createPinia()
pinia.use(piniaPersistedState)


createApp(App)
    .use(router)
    .use(pinia)
    .mount("#app")

router.push("/home")