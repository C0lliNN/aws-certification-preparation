import logo from './logo.svg';
import './App.css';
import { useEffect, useState } from 'react';
import { getPerson } from './service/api';

function App() {
  const [name, setName] = useState(null);
  const [age, setAge] = useState(0);

  async function updatePerson() {
    const newPerson = await getPerson();
    setName(newPerson.name);
    setAge(newPerson.age);
  }

  useEffect(() => {
    updatePerson();
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <p>
          Name: {name}, Age: {age}
        </p>
        <button onClick={updatePerson}>New person</button>
      </header>
    </div>
  );
}

export default App;
