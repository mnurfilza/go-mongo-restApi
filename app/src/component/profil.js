import React, { Component } from 'react'
import {Container,Row,Col} from 'reactstrap'
import bgImage from '../images/our.JPG'
import './css/profil.css'

export default class Profil extends Component {
    render() {
        return (
            <div>
                <Container fluid="sm">
                    <Row>
                        <Col>
                            <div className="title">
                                <h1 className="display-4">About Us</h1>
                            </div>
                            <hr className="line" style={{border:"5px solid #757575", borderRadius:"10px"}}/>
                            <br/>
                            <br/>
                        </Col>
                    </Row>
                    <Row>
                        <Col m="12">
                            <Container>
                                <Row>
                                    <Col sm="6">
                                        <img src={bgImage} alt="Smiley face" width="60%" style={{objectFit:"scale-down"}}/>
                                    </Col>

                                    <Col sm="6">
                                        <Container>
                                            <Row>
                                                <Col sm="12">
                                                    <p className="display-4">Rahmi Fauziah S.Psi</p>
                                                </Col>
                                            </Row>
                                            <Row>
                                                <h6 className="bpkIbuCpw">Putri Dari Bapak Ibrahim & Ibu Iva Herlistriawati</h6>
                                            </Row>
                                               <Row>
                                                <Col sm="12">
                                                    <h2 className="display-2">
                                                        &    
                                                    </h2>
                                                </Col>
                                               </Row> 
                                            <Row>
                                                <Col sm="12">
                                                    <p className="display-4">M. Nurfilza Abdul Choir S.Kom</p>
                                                </Col>
                                             </Row>
                                             <Row>
                                                 <h6 className="bpkIbuCpp">Putri Dari Bapak Kosasih & Ibu Wiji Lestari</h6>
                                            </Row>
                                        </Container>
                                    </Col>
                                    
                                </Row>
                            </Container>
                        </Col>

                       
                    </Row>
                </Container>
            </div>
        )
    }
}
