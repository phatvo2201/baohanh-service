import React from "react";

function App() {
  const callApi = async () => {
    const res = await fetch("http://localhost:8080/api/movies/", {
      headers: { "Authorization": "Bearer test" } // if your API requires auth
    });
    const data = await res.json();
    alert(JSON.stringify(data, null, 2));
  };

  return (
    <div>
      <h1>Test CORS with Go API</h1>
      <button onClick={callApi}>Call /api/movies</button>
    </div>
  );
}

export default App;