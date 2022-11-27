<template>
  <v-form>
    <v-row>
      <v-col cols="12" sm="6">
        <v-text-field
          v-model="bookmarkURL"
          label="Bookmark URL"
          outlined
          dense
        ></v-text-field>
      </v-col>
      <v-col cols="12" sm="6">
        <v-btn
          :disabled="!bookmarkURL"
          color="primary"
          class="mr-4"
          :loading="loading"
          @click="submit"
        >
          Submit
        </v-btn>
      </v-col>
    </v-row>
  </v-form>
</template>

<script>
export default {
  name: 'BookmarkSubmissionForm',
  props: {
    loading: Boolean,
  },
  data: () => ({
    bookmarkURL: '',
  }),
  computed: {},
  methods: {
    submit() {
      this.$emit('submit', this.bookmarkURL)
      this.bookmarkURL = ''
    },
  },
}
</script>
