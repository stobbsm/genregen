<template>
  <v-app>
    <v-main>
      <GetGenre/>
    </v-main>
  </v-app>
</template>

<script>
import GetGenre from './components/genre.vue'

export default {
  name: 'App',

  components: {
    GetGenre,
  },

  data: () => ({
    //
  }),
}
</script>
