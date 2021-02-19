export default {
  setupWebSocketConn(callback) {
    let ws = new WebSocket(`ws://booky-env-1.us-east-1.elasticbeanstalk.com:5000/api/connect`)
    ws.onopen = function () {
      console.log('Connected.')
    }

    ws.onmessage = function (message) {
      if (!message.data) {
        return
      }

      let data = JSON.parse(message.data)
      callback(data)
    }
    ws.onerror = function (event) {
      console.error('WebSocket error observed:', event)
    }
  }
}
