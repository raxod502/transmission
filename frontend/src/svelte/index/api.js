export class API {
  constructor({ onStateUpdate }) {
    if (!onStateUpdate) throw new Error("API: onStateUpdate is required");
    this.socket = null;
    this.onStateUpdate = onStateUpdate;
  }
  connect() {
    if (this.socket) {
      this.socket.close();
    }
    console.log("Connecting to websocket...");
    this.socket = new WebSocket(
      (document.location.protocol === "http:" ? "ws://" : "wss://") +
        document.location.host +
        "/api/v1/ws"
    );
    window.socket = this.socket;
    this.socket.addEventListener("open", () => {
      console.log("Connected to websocket.");
    });
    this.socket.addEventListener("message", (event) => {
      let message;
      try {
        message = JSON.parse(event.data);
      } catch (err) {
        console.error("Ignoring malformed JSON in server message:", event.data);
        return;
      }
      // The only message we support is the server sending us the
      // entire state.
      this.onStateUpdate(message);
    });
  }
}
