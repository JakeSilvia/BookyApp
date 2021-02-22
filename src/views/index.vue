<template>
  <div class="section mt-5">
    <h3 class="title is-3">Books</h3>
    <b-loading :active.sync="isLoading"></b-loading>
    <div class="has-text-right mb-5 is-fullwidth-mobile">
      <button class="button is-primary is-small" @click="openNewBookModal">
        <b-icon icon="plus"/>
      </button>
    </div>
    <div class="container">
      <div v-if="books.length > 0">
        <b-table
            :data="books"
            detailed
            detail-key="isbn"
            default-sort="Title"
            show-detail-icon
            striped>
          <b-table-column
              v-slot="props"
              field="isbn"
              sortable
              searchable
              label="ISBN">
            <span v-if="props.row.isbn === ''" class="has-text-grey-light">N/A</span>
            <span v-else>{{ props.row.isbn }}</span>
          </b-table-column>
          <b-table-column
              v-slot="props"
              field="title"
              sortable
              searchable
              label="Title">
            <span v-if="props.row.title === ''" class="has-text-grey-light">N/A</span>
            <span v-else>{{ formatName(props.row.title) }}</span>
          </b-table-column>
          <b-table-column
              v-slot="props"
              field="author"
              sortable
              searchable
              label="Author">
            <span v-if="props.row.author === ''" class="has-text-grey-light">N/A</span>
            <span v-else>{{ formatName(props.row.author) }}</span>
          </b-table-column>
          <b-table-column
              v-slot="props"
              field="status"
              sortable
              centered
              label="Status">
            <span v-show="props.row.status === 'Available'" class="has-text-success">Available</span>
            <span v-show="props.row.status === 'Checked Out'" class="has-text-danger">Unavailable</span>
          </b-table-column>
          <b-table-column
              v-slot="props"
              centered
              label="Actions">
            <b-dropdown @input="changeBook($event, props.row, props.index)">
              <template #trigger="{ active }">
                <button class="button has-text-primary is-text no-underline">
                  Actions
                  <b-icon class="ml-1" :icon="active ? 'menu-up' : 'menu-down'"/>
                </button>
              </template>

              <b-dropdown-item :value="editEventType">{{ editEventType }}</b-dropdown-item>
              <b-dropdown-item :value="deleteEventType">{{ deleteEventType }}</b-dropdown-item>
              <b-dropdown-item v-if="props.row.status === 'Available'" :value="checkOutEventType">{{ checkOutEventType }}</b-dropdown-item>
            </b-dropdown>
          </b-table-column>
          <template #detail="props">
            <span v-html="props.row.description"></span>
            <span v-if="props.row.description === ''"
                  class="has-text-grey-light"
            >No book description. Please add one by clicking the Edit action</span>
          </template>
        </b-table>
      </div>
      <div v-if="books.length === 0" class="has-background-white-ter has-text-grey">
        <h4 class="is-4 subtitle has-text-grey-light has-text-weight-light ">No Books have been created. Please
          create a new one in the top right</h4>
      </div>
      <b-modal :active.sync="editModalActive" width="640" has-modal-card>
        <book-edit-modal @update-success="refreshBook" :book="currentBook"></book-edit-modal>
      </b-modal>
      <b-modal :active.sync="newModalActive" width="640" has-modal-card>
        <book-create-modal @create-success="addBook"></book-create-modal>
      </b-modal>
    </div>
  </div>
</template>

<script>
import BookEditModal from '@components/book-edit-modal'
import BookCreateModal from '@components/book-create-modal'


export default {
  name: "books",
  components: {BookEditModal, BookCreateModal},
  data: function () {
    return {
      maxCharacterLength: 30,
      editEventType: 'Edit',
      deleteEventType: 'Delete',
      checkOutEventType: 'Check Out',
      isLoading: false,
      books: [],
      editModalActive: false,
      newModalActive: false,
      deleteModalActive: false,
      currentBook: null,
      currentBookIndex: null
    }
  },
  beforeMount() {
    this.getBooks()
    this.$listenForEvents(this.handleMessage)
  },
  methods: {
    handleMessage(event) {
      console.log("Event:", event)
      switch (event.action){
        case "updated":
          this.$notification('is-white', `<b>${this.formatName(event.book_item.title)}</b> has been updated`)
          this.findAndUpdateBook(event.book_item)
          break
        case "inserted":
          this.$notification('is-white', `<b>${this.formatName(event.book_item.title)}</b> has been created`)
          this.addBook(event)
          break
        case "deleted":
          this.findAndDeleteBook(event.book_item)
          break
      }
    },
    findAndDeleteBook(book) {
      for (let i = 0; i < this.books.length; i++) {
        if (this.books[i].isbn === book.isbn) {
          this.books.splice(i, 1)
          return
        }
      }
    },
    findAndUpdateBook(book) {
      for (let i = 0; i < this.books.length; i++) {
        if (this.books[i].isbn === book.isbn) {
          this.$set(this.books, i, book)
          return
        }
      }
      this.books.push(book)
    },
    formatName(name) {
      if (!name || name.length === 0) {
        return ''
      }
       return name.length > this.maxCharacterLength ? name.substr(0, this.maxCharacterLength) + '...' : name
    },
    getBooks() {
      this.isLoading = true
      this.$ajax.getBooks().then(resp => {
        if (resp.err != null) {
          this.$notification('is-danger', resp.err)
          return
        }

        if (!resp.data) {
          return
        }
        this.books = resp.data
      }).finally(_ => {
        this.isLoading = false
      })
    },
    refreshBook(updatedBook) {
      this.$set(this.books, this.currentBookIndex, updatedBook)
      this.editModalActive = false
    },
    addBook(newBook) {
      this.books.push(newBook)
      this.newModalActive = false
    },
    openNewBookModal() {
      this.newModalActive = true
    },
    openEditModal(book, index) {
      this.currentBook = book
      this.currentBookIndex = index
      this.editModalActive = true
    },
    openCheckOutBookDialog(book, index) {
      this.$buefy.dialog.confirm({
        title: 'Check Out',
        message: `Would you like to check out <b>${book.title}</b> ?`,
        confirmText: 'Yes',
        type: 'is-success',
        onConfirm: () => {
          this.checkoutBook(book, index)
        }
      })
    },
    checkoutBook(book, index) {
      book.status = "Checked Out"
      this.$ajax.updateBook(book).then(resp => {
        if (resp.err != null) {
          this.$notification('is-danger', `Could not delete ${book.title}: `, resp.err)
          return
        }

        this.$set(this.books, index, book)
        this.$notification('is-success', `<b>${this.formatName(book.title)}</b> has been checked out`)
      })
    },
    openDeleteBookDialog(book, index) {
      this.$buefy.dialog.confirm({
        title: 'Delete Book',
        message: `Are you sure you want to delete <b>${book.title}<b></b>`,
        confirmText: 'Delete',
        type: 'is-danger',
        onConfirm: () => {
          this.deleteBook(book, index)
        }
      })
    },
    deleteBook(book, index) {
      this.$ajax.deleteBook(book.isbn).then(resp => {
        if (resp.err != null) {
          this.$notification('is-danger', `Could not delete ${book.title}: `, resp.err)
          return
        }

        this.$notification('is-success', book.title + ' deleted')
        this.books.splice(index, 1)
      })
    },
    changeBook(event, book, index) {
      switch (event) {
        case this.editEventType:
          this.openEditModal(book, index)
          break
        case this.deleteEventType:
          this.openDeleteBookDialog(book, index)
          break
        case this.checkOutEventType:
          this.openCheckOutBookDialog(book)
          break
      }
    }
  }
}
</script>

<style scoped lang="scss">
.no-underline {
  text-decoration: none !important;
}

.has-columns-vertically-aligned-middle td {
  vertical-align: middle;
}


</style>
