import React from "react"
import {Card,} from "react-bootstrap"

const Cards = () => {
  return (
    <>
      <section className='about topMarign'>
        <div className='heading'>
        <div className="cards">
            <center><h1>Choose a Course!</h1> </center>s
          <div className='container flex d-flex gap-4 mt-5'>
<Card style={{ width: '18rem' }} className="border-0  bg-transparent">
<Card.Body className="card-body1">
<center><img className="card-img1" src="./assets/vocab.png" /></center>
  <Card.Title style={{color : "white"}}><center>VOCAB</center></Card.Title>
  <Card.Text style={{color : "white", textAlign:"justify"}}>
    Some quick example text to build on the card title and make up the bulk of
    the card's content.
  </Card.Text>
  <center><button className="start1">Start</button></center>
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
  <center><button className="start2">Start</button></center>
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
  <center><button className="start3">Start</button></center>
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
