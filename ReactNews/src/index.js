import React from "react";
import ReactDOM from "react-dom";

import "./styles.css";
import { EventsSection } from "./EventsSection";

function App() {
  return (
    <div className="App">
      <EventsSection />
    </div>
  );
}

const rootElement = document.getElementById("root");
ReactDOM.render(<App />, rootElement);
