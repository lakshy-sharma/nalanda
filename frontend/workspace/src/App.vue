<template>
  <div>
    <Header />
    <div>
      <router-view/>
    </div>
    <Footer />
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";
import store from "./components/store";

const getCookie = (name) => {
  return document.cookie.split("; ").reduce((r, v)=> {
    const parts = v.split("=");
    return parts[0] === name ? decodeURIComponent(parts[1]) : r;
  }, "")
}

export default {
  name: "App",
  components: {
    Header,
    Footer,
  },
  data() {
    return {
      store
    }
  },
  beforeMount() {
    let data = getCookie("_site_data")
    if (data !== "") {
      let cookieData = JSON.parse(data);
      // update the store.

      store.token = cookieData.toke.token;
      store.user = {
        id: cookieData.user.id,
        first_name: cookieData.user.first_name
      }
    }
  }
}
</script>

<style>
</style>