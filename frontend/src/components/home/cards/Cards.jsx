import {Card,} from "react-bootstrap"
import "./style.css"
import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import { Link } from "react-router-dom"

const Cards = () => {
  const [show, setShow] = useState(false);
  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);
  return (
    <>
      <section className='about topMarign'>
        <div className='heading'>
        <div className="cards" id="course">
            <center><h1>Choose a Course!</h1> </center>
          <div className='container flex d-flex gap-4 mt-5'>
<Card style={{ width: '18rem' }} className="border-0  bg-transparent">
<Card.Body className="card-body1">
<center><img className="card-img1" src="./assets/vocab.png" /></center>
  <Card.Title style={{color : "white"}}><center>VOCAB</center></Card.Title>
  <Card.Text style={{color : "white", textAlign:"justify"}}>
    Some quick example text to build on the card title and make up the bulk of
    the card's content.
  </Card.Text>
  <center><button className="start1" onClick={handleShow}>
        Start
      </button></center>
      <Modal show={show} onHide={handleClose}>
        <Modal.Body>We see you haven't logged in yet, would you like to do it now?
        <Button className="login2" variant="warning" onClick={handleClose}>
        <Link to='login' style={{textDecoration: "none", color : "black"}}>Log In</Link>
          </Button>
        </Modal.Body>
      </Modal>
</Card.Body>
</Card>

<Card style={{ width: '18rem' }} className="border-0 bg-transparent">
<Card.Body className="card-body2">
<center><img className="card-img2" src="./assets/grammar.png" /></center>
  <Card.Title style={{color : "white"}}><center>GRAMMAR</center></Card.Title>
  <Card.Text style={{color : "white", textAlign:"justify"}}>
    Some quick example text to build on the card title and make up the bulk of
    the card's content.

  </Card.Text>
  <center><button className="start2" onClick={handleShow}>
        Start
      </button></center>
</Card.Body>
</Card>

<Card style={{ width: '18rem' }} className="border-0  bg-transparent">
<Card.Body className="card-body3">
<center><img className="card-img3" src="./assets/tenses.png" /></center>
  <Card.Title style={{color : "white"}}><center>TENSES</center></Card.Title>
  <Card.Text style={{color : "white", textAlign:"justify"}}>
    Some quick example text to build on the card title and make up the bulk of
    the card's content.

  </Card.Text>
  <center><button className="start3" onClick={handleShow}>
        Start
      </button></center>
</Card.Body>
</Card>
</div>
</div>  
        </div>
      </section>

      <div className="boxCard"></div>
      
    </>
  )
}

export default Cards
