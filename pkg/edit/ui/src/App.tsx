import React from "react";
import { Route, BrowserRouter as Router, Switch } from "react-router-dom";

import Main from "./views/Main";

const App = () => (
  <Router>
    <Switch>
      <Route path="/">
        <Main />
      </Route>
    </Switch>
  </Router>
);

export default App;
