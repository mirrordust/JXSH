import React from "react";


class App extends React.Component {
  render(): JSX.Element {
    return (
      <h1>
        My React and TypeScript App!!{" "}
        {new Date().toLocaleString()}
      </h1>
    );
  }
}


export default App;
