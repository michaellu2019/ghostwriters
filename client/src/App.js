import { useState, useEffect } from 'react';
import { Route, Link } from 'react-router-dom';

import Bookshelf from './pages/Bookshelf';
import Story from './pages/Story';

import './App.css';

function App() {
  return (
    <div className="wrapper">
      <header className="primary">
          <nav>
            <span className="heading"><Link className="nav-link" to="/">Ghostwriters</Link></span>

            <ul className="nav-buttons links">
              <li><Link className="nav-link" to="/">Home</Link></li>
            </ul>
          </nav>
        </header>

      <main className="primary">
        <Route exact path="/" component={Bookshelf} />
        <Route exact path="/story/:id" component={Story} />
      </main>
    </div>
  );
}

export default App;
