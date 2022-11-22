<template>
  <div>
    <span v-if="isLoggedIn">
      <v-btn color="blue-grey" class="white--text" to="/account" nuxt>
        <v-icon> mdi-account </v-icon>
      </v-btn>
    </span>
    <span v-else>
      <v-btn color="blue-grey" class="white--text" @click="login">
        <v-icon left> mdi-login </v-icon>
        <span> Login </span>
      </v-btn>
    </span>
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
export default {
  name: 'AccountButton',
  computed: {
    ...mapGetters({
      isLoggedIn: 'auth/isLoggedIn',
    }),
  },
  methods: {
    login() {
      this.$store.dispatch('auth/login')
    },
  },
}
</script>
