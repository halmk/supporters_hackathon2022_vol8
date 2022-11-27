<template>
  <v-row justify="center">
    <v-col cols="12" align="center">
      <img :src="userPhoto" />
    </v-col>

    <v-col cols="12" align="center">
      <div class="text-h3">{{ userName }}</div>
    </v-col>
    <v-col cols="12" align="center">
      <div>
        <v-btn color="orange" class="white--text" @click="logout">
          <v-icon> mdi-logout-variant </v-icon>
          <v-list-item-title> ログアウト </v-list-item-title>
        </v-btn>
      </div>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'AccountPage',
  data: () => ({
    userName: '',
    userPhoto: '',
  }),
  created: function () {
    this.userName = this.$store.getters['auth/getUserName']
    this.userPhoto = this.$store.getters['auth/getUserPhoto']
  },
  methods: {
    logout: function () {
      localStorage.removeItem('name')
      localStorage.removeItem('uuid')
      localStorage.removeItem('photo')
      this.$store.dispatch('auth/logout')
      this.$router.push('/')
    },
  },
}
</script>
