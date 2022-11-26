<template>
  <v-form v-model="valid">
    <v-row>
      <v-col cols="12" sm="6">
        <v-text-field
          v-model="categoryName"
          :rules="nameRules"
          label="Category Name"
          required
          dense
        ></v-text-field>
      </v-col>
      <v-col cols="12" sm="2">
        <v-btn :disabled="!valid" color="primary" class="mr-4" @click="submit">
          Add
        </v-btn>
      </v-col>
    </v-row>
  </v-form>
</template>
<script>
export default {
  name: 'CategoryCreateForm',
  data: () => ({
    categoryName: '',
    valid: false,
    nameRules: [(v) => !!v || 'required'],
  }),
  computed: {},
  methods: {
    submit() {
      this.$emit('submit', this.categoryName)
    },
  },
}
</script>
