<template>
  <div>
    <!--      <b-loading :active.sync="Vue.prototype.$devicesLoading" is-full-page/>-->
    <b-navbar
        fixed-top
        class="is-primary navbar is-spaced">
      <template slot="brand">
      </template>
      <template slot="start" class="has-text-centered ">
        <b-navbar-item
            v-for="route in routes"
            tag="router-link"
            class="has-text-weight-bold"
            :to="route.path">
          {{ route.label }}
        </b-navbar-item>
        <div class="navbar-burger"/>
      </template>
      <template slot="end">
        <a class="has-text-white" @click="downloadReport()">Export <b-icon icon="arrow-collapse-down" size="is-small"/> </a>
      </template>
    </b-navbar>
  </div>
</template>

<script>
export default {
  name: 'BookNav',
  data() {
    return {
      routes: [
        {
          label: "Home",
          path: "/"
        },
        {
          label: "Events",
          path: "/events"
        },
      ]
    }
  },
  methods: {
    downloadReport() {
      window.open("/api/books/report", '_blank')
    }
  }
}
</script>

<style scoped>
.navbar .navbar-brand {
  text-align: center;
  display: block;
  width: 100%;
}

.navbar .navbar-brand > .navbar-item,
.navbar .navbar-brand .navbar-link {
  display: inline-block;
}
</style>
