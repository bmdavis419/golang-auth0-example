import { useState } from "react";
import reactLogo from "./assets/react.svg";
import "./App.css";
import { useAuth0 } from "@auth0/auth0-react";

function App() {
  const [count, setCount] = useState(0);

  const { loginWithRedirect, isAuthenticated, logout, getAccessTokenSilently } =
    useAuth0();

  const showApiCall = () => {
    getAccessTokenSilently().then((token) => {
      fetch("http://localhost:8080/user/me", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
        .then((response) => response.json())
        .then((data) => alert(data.message));
    });
  };

  return (
    <div className="App">
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src="/vite.svg" className="logo" alt="Vite logo" />
        </a>
        <a href="https://reactjs.org" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>

      {isAuthenticated ? (
        <>
          <button
            onClick={() =>
              logout({ logoutParams: { returnTo: window.location.origin } })
            }
          >
            logout
          </button>
          <button onClick={() => showApiCall()}>show api call</button>
        </>
      ) : (
        <button onClick={() => loginWithRedirect()}>login</button>
      )}
    </div>
  );
}

export default App;
