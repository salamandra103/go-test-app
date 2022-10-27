import React, { useEffect } from 'react';

import {A} from './A'

function App() {

  useEffect(() => {
    fetch("http://localhost:8080/hi", {
      method: "GET"
    }).then(res => res.json()).then(res => {

      console.log(res);
      // debugger;
    })
  }, [])

  return (
    <div className="App">
      12313ввы
      {/* <A/> */}
    </div>
  );
}

export default App;
