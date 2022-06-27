import React from "react"
import "./style.css"

const Footer = () => {
  return (
    <>
      <footer>
        <div className='container grid1'>
          <div className='box'>
            <img src='./assets/logo3.png' alt='' />
            <p>Edumar also known as Educational Grammar
  is a website-based english quiz where there are three types of courses namely Vocab, Grammar and Tenses.</p>
            <br></br>
            <div className='SocailIcon'>
              <i className='fab fa-facebook-f facebook'></i>
              <i className='fab fa-instagram instagram'></i>
              <i className='fab fa-twitter twitter'></i>
              <i className='fab fa-youtube youtube'></i>
            </div>
          </div>

          <div className='box'>
            <h2>Course Selection</h2>
            <ul>
              <li>Vocab</li>
              <li>Grammar</li>
              <li>Tenses</li>
            </ul>
          </div>

          <div className='box'>
            <h2>Help & Support</h2>
          <ul>
              <li>Home</li>
              <li>About Us</li>
              <li>Scoreboard</li>
            </ul>
          </div>

          <div className='box'>
            <h2>Get in Touch</h2>
            <div className='icon'>
              <i class='fa-solid fa-location-dot'></i>
              <label>DKI Jakarta, Indonesia</label>
            </div>
            <div className='icon'>
              <i class='fa fa-phone'></i>
              <label>Phone : +62 858-6030-1670</label>
            </div>
            <div className='icon'>
              <i class='fa fa-envelope'></i>
              <label>Email : support@edumar.com</label>
            </div>
          </div>
        </div>
        <div className='d-flex justify-content-center legal container'>
          <p className="copy" style={{textAlign:"center"}}>Copyright @2022. All rights reserved.</p>
        </div>
        <a href="#" class="up-icon"><i class="fas fa-chevron-circle-up"></i></a>
      </footer>
    </>
  )
}

export default Footer
