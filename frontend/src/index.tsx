import React from 'react';
import ReactDOM from 'react-dom';
import reportWebVitals from './reportWebVitals';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom';

import HomePage from './pages/home';
import AdminPage from './pages/admin';
import SignIn from './pages/login';

function App() {
  return (
    <div className="App container">
      <div className="jumbotron">
        <Router>
          <Switch>
            <Route exact path="/" component={HomePage}/>
            <Route path="/admin" component={AdminPage}/>
            <Route path="/login" component={SignIn}/>
          </Switch>
        </Router>
      </div>
    </div>
  );
}
const rootElement = document.getElementById('root');
ReactDOM.render(<App/>, rootElement);


// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
