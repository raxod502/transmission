export class API {
  constructor() {
    this.socket = null;
  }
  connect(onStateUpdate) {
    if (this.socket) {
      this.socket.close();
    }
    console.log("Connecting to websocket...");
    this.socket = new WebSocket(
      (document.location.protocol === "http:" ? "ws://" : "wss://") +
        document.location.host +
        "/api/v1/ws"
    );
    this.socket.addEventListener("open", () => {
      console.log("Connected to websocket.");
    });
    this.socket.addEventListener("message", (event) => {
      console.log(event);
    });
  }
}
