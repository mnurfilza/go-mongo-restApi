
import {Jumbotron,Container,Button} from 'reactstrap'

import React, { Component } from 'react'
import Count from './countDown'
import './css/jumbotron.css'
import bgImage from '../images/jadi.jpg'
export default class Jumbo extends Component {

    constructor(props) {
        super(props);
        this.handleClick = this.handleClick.bind(this)
    }
    handleClick(){
        console.log("ini diklick")
    }
    
  render() {
    const currentDate = new Date();
    const year = currentDate.getFullYear()+1;
    return (
      <div>
      <Jumbotron fluid style={{ 
          backgroundImage: `url(${bgImage})`, 
          backgroundSize: 'cover',
          height:'95vh',
          width:'100%',
          opacity:'200',
          backgroundPosition:'35%',
          position:'relative',
          color:'white',
          backgroundRepeat:'no-repeat',
        backgroundAttachment:'fixed'}}>

      <Container fluid style={{paddingTop:'10%'}}>
        <h2 className="display-4">Heppy & Bima</h2>
        <p className="lead">We are getting Married</p>
        <Count date={`${year}-08-19T00:00:00`}/>
        <Button color="secondary" onClick={this.handleClick} style={{borderRadius:'25px', padding:'10px 15px 10px 15px'}}>Save The Date</Button>
      </Container>

        
    </Jumbotron>
      </div>
    )
  }
}
