import React, { useEffect } from 'react';

function App() {

  useEffect(() => {
    fetch("http://localhost:8080/hi", {
      method: "GET"
    }).then(res => res.json()).then(res => {

      console.log(res);
      debugger;
    })
  }, [])

  return (
    <div className="App">
      Test
    </div>
  );
}

export default App;
