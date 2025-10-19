import { useState } from "react";
import "./App.css";

function App() {
  const [url, setURL] = useState("");
  const [redirect, setRedirect] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    const response = await fetch("http://localhost:8080", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ url: url }),
    });

    if (!response.ok) return;

    const result = await response.json();
    setRedirect(result.url);
  };

  return (
    <main>
      <h1>URL Shortner</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={url}
          onChange={(e) => setURL(e.target.value)}
        />

        <button type="submit">Generate URL</button>
      </form>

      {redirect && <div>
        <p>Successfully generated your short link</p>
        <a href={redirect}>{redirect}</a>
      </div>}
    </main>
  );
}

export default App;
