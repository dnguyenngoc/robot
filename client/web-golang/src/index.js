import React from "react";
import { BrowserRouter as Router, Switch, Redirect, Route } from "react-router-dom";
import ReactDOM from "react-dom"

import RoutesConfig from "./routes/RoutesConfig";
import AppRoute from "./routes/AppRoute";


export class App extends React.Component {
  render() {
    return ( 
      <Router>
         <div>
          <Switch>
          {
            RoutesConfig.map(
              route =>(
                <AppRoute 
                  key={route.path}
                  path={route.path}
                  container={route.container}
                  auth={route.isPrivate}
                /> 
              )
            )
          }
            <Route exact path="/" render={() => (<Redirect to="/login"/>)} />
          </Switch>
         </div>
      </Router>
    );
  }
}


ReactDOM.render( 
  <App/> ,
  document.getElementById("root")
);