import { useState, useEffect } from 'react';
import { Route, Link } from 'react-router-dom';

import Bookshelf from './pages/Bookshelf';
import Story from './pages/Story';

import './App.css';

function App() {
  const [loggedIn, setLoggedIn] = useState(false);
  const [user, setUser] = useState({});

  function login(user) {
    if (user.hasOwnProperty("userId") && user.hasOwnProperty("username")) {
      setLoggedIn(true);
      setUser(user);
      console.log(user);
    }
  }

  return (
    <div className="wrapper">
      <header className="primary">
          <nav>
            <span className="heading"><Link className="nav-link" to="/">Ghostwriters</Link></span>

            <ul className="nav-buttons links">
              <li><Link className="nav-link" to="/">Stories</Link></li>
            </ul>
          </nav>
        </header>

      <main className="primary">
        <Route exact path="/" render={() => <Bookshelf loggedIn={loggedIn} user={user} login={login} />} />
        <Route exact path="/story/:id" render={() => <Story loggedIn={loggedIn} user={user} login={login} />} />
      </main>
    </div>
  );
}

export default App;
