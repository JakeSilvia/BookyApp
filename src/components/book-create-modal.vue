<template>
  <div class="modal-card">
    <div class="modal-card-head">
      <p class="title is-4">New Book</p>
    </div>
    <div class="modal-card-body">
      <div class="mt-3">
        <b-field grouped>
          <b-field label="ISBN">
            <b-input v-model="stagedChanges.isbn" type="number"></b-input>
          </b-field>
          <b-field label="Title">
            <b-input v-model="stagedChanges.title" type="text"></b-input>
          </b-field>
          <b-field label="Author">
            <b-input v-model="stagedChanges.author" type="text"></b-input>
          </b-field>
        </b-field>
      </div>
      <div class="mt-5">
        <label class="label">Description</label>
        <vue-editor v-model="stagedChanges.description"></vue-editor>
      </div>
      <div class="mt-5">
        <b-tooltip position="is-right" size="is-small" multilined label="This will override whoever currently has the book checked out.">
          <label class="label">Status<b-icon size="is-small" icon="help-circle"></b-icon></label>
        </b-tooltip>
        <b-select v-model="stagedChanges.status">
            <option value="Checked Out">Checked Out</option>
            <option value="Available">Available</option>
        </b-select>
      </div>
    </div>
    <div class="modal-card-foot">
      <button class="button is-success is-fullwidth" @click="update()">Save</button>
    </div>
  </div>
</template>

<script>
import {VueEditor} from "vue2-editor";

export default {
  name: "book-card",
  components: {VueEditor},
  data() {
    return {
      stagedChanges: {
        status: "Available",
      }
    }
  },
  created() {},
  methods: {
    update() {
      this.$ajax.createBook(this.stagedChanges).then(resp => {
        if (resp.err != null) {
          this.$notification('is-danger', resp.err)
          return
        }

        this.$emit('create-success', this.stagedChanges)
      })
    }
  }
}
</script>

<style scoped>

</style>
