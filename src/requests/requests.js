import axios from 'axios'

export default {
  catchErrors (error) {
    if (error.response) {
      return { data: null, err: error.response.data }
    } else if (error.request) {
      return { data: null, err: error.request }
    } else {
      return { data: null, err: error.message }
    }
  },
  getEvents (count) {
    return axios.get('/api/events?count=' + count, {}).then(resp => {
      return { data: resp.data, err: null }
    }).catch(error => this.catchErrors(error))
  },
  getBooks () {
    return axios.get('/api/books', {}).then(resp => {
      return { data: resp.data, err: null }
    }).catch(error => this.catchErrors(error))
  },
  getReport () {
    return axios.get('/api/books/report', {}).then(resp => {
      return { data: resp.data, err: null }
    }).catch(error => this.catchErrors(error))
  },
  updateBook (book) {
    return axios.post('/api/books', book).then(resp => {
      return { data: resp.data, err: null }
    }).catch(error => this.catchErrors(error))
  },
  createBook (book) {
    return axios.put('/api/books', book).then(resp => {
      return { data: resp.data, err: null }
    }).catch(error => this.catchErrors(error))
  },
  deleteBook (isbn) {
    return axios.delete('/api/books?isbn=' + isbn, {}).then(resp => {
      return { data: resp.data, err: null }
    }).catch(error => this.catchErrors(error))
  }
}
