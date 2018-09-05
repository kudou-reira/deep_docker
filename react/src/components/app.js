import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Switch, Redirect, Link } from 'react-router-dom';

import HomeScreen from '../screens/homeScreen';
import AltScreen from '../screens/altScreen';

// control the routing here

class App extends Component {
  render() {
    return (
      <Router>
        <div>
          <Switch>
            <Route exact path="/" component={HomeScreen} />
            <Route exact path="/alt" component={AltScreen} />
          </Switch>
        </div>
      </Router>
    );
  }
}

export default App;
