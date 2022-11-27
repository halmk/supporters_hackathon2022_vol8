<template>
  <v-form>
    <v-row>
      <v-col cols="12" sm="6">
        <v-text-field
          v-model="categoryName"
          label="Category Name"
          outlined
          dense
        ></v-text-field>
      </v-col>
      <v-col cols="12" sm="2">
        <v-btn
          :disabled="!categoryName"
          color="primary"
          class="mr-4"
          @click="submit"
        >
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
  }),
  computed: {},
  methods: {
    submit() {
      this.$emit('submit', this.categoryName)
    },
  },
}
</script>
