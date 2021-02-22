<template>
<div class="section">
  <h3 class="title is-3">Events</h3>
  <div class="container">
    <div v-if="events.length > 0">
      <b-table
          :data="events"
          detailed
          detail-key="date"
          default-sort="Title"
          show-detail-icon
          striped>
        <b-table-column
            v-slot="props"
            field="book_item.isbn"
            searchable
            label="ISBN">
          {{ props.row.book_item.isbn }}
        </b-table-column>
        <b-table-column
            v-slot="props"
            searchable
            label="Action">
          {{ props.row.action }}
        </b-table-column>
        <b-table-column
            v-slot="props"
            searchable
            label="Date">
          <pre>{{ props.row.date.toString() }}</pre>
        </b-table-column>
        <template #detail="props">
          <div class="container">
            <pre>{{ props.row.book_item }}</pre>
          </div>
        </template>
      </b-table>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: "events",
  data() {
    return {
       events: []
    }
  },
  beforeMount() {
    this.$listenForEvents(this.handleEvent)
    this.getEvents()
  },
  methods: {
    handleEvent(event) {
      this.events.unshift(event)
    },
    getEvents() {
      this.$ajax.getEvents(50).then(resp => {
        if (resp.err != null) {
          this.$notification('is-danger', resp.err)
          return
        }
        console.log("E: ", resp.data)
        this.events = resp.data
      })
    }
  }
}
</script>

<style scoped>
</style>
