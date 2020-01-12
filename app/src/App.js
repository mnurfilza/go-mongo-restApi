
import React, { Component } from 'react';
import Jumbotron from './component/jumbotron'
import Profil from './component/profil'
import './App.css';

class App extends Component{
  render(){
    
    return (
      <div className="App">
        <Jumbotron></Jumbotron>
        <Profil></Profil>
      </div>
    );
  }
}

export default App;
