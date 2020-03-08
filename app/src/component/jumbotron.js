
import {
  Jumbotron,
  Container,
  Button,
 
} from 'reactstrap'

import React, { Component } from 'react'
import Count from './countDown'
import './css/jumbotron.css'
import bgImage from '../images/FER26360.JPG'
export default class Jumbo extends Component {

    constructor(props) {
        super(props);
        this.handleClick = this.handleClick.bind(this)
     
    }
    handleClick(){
        console.log("ini diklick")
        console.log("ss");
        
    }

  

    
  render() {
    const currentDate = new Date();
    const year = currentDate.getFullYear();
    return (
      <div>
      <Jumbotron fluid style={{ 
          backgroundImage: `url(${bgImage})`, 
          backgroundSize: 'cover',
          height:'100vh',
          width:'100%',
          opacity:'200',
          backgroundPosition:'50% 50% ',
          position:'relative',
          color:'white',
          backgroundRepeat:'no-repeat',
        backgroundAttachment:'fixed'}}>

        
          <div className="content">
          <Container fluid="md">
            <h1 className="display-2">Rahmi & Filza</h1>
            <p className="lead">We are getting Married</p>
            <Count date={`${year}-04-04T00:00:00`}/>
            <Button color="secondary" onClick={this.handleClick} style={{borderRadius:'15px',opacity:'0.7', marginTop:'50px',padding:'10px 15px 10px 15px'}}>Save The Date</Button>
          </Container>
        </div>
        
      </Jumbotron>
      </div>
    )
  }
}
