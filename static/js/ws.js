const webSocketConnection = () => {
   // Connect to the WebSocket server
   const socket = new WebSocket('ws:///ws');

   // Display connection status
   socket.onopen = () => {
     console.log("Connected to server.");
     document.getElementById("messages").innerText = "Connected to server!";
   };

   // Display any received messages from the server
   socket.onmessage = (event) => {
     console.log("Message from server:", event.data);
     const messagesDiv = document.getElementById("messages");
     messagesDiv.innerHTML += `<br>Server: ${event.data}`;
   };

   // Handle connection errors
   socket.onerror = (error) => {
     console.error("WebSocket Error: ", error);
     document.getElementById("messages").innerText = "Error connecting to server.";
   };

   // Handle WebSocket close
   socket.onclose = () => {
     console.log("Disconnected from server.");
     document.getElementById("messages").innerText = "Disconnected from server.";
   };

}

// webSocketConnection()
