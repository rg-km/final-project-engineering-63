import React, { useState } from "react"
import {NavDropdown} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import { Link } from "react-router-dom"
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import LogoutIcon from '@mui/icons-material/Logout';

const Header = () => {
  const [sidebar, setSidebar] = useState(false)
  window.addEventListener("scroll", function () {
    const header = document.querySelector(".header")
    header.classList.toggle("active", window.scrollY > 180)
  })
  return (
    <>
      <header className='header'>
        <div className='container flex'>
          <div className='logo'>
            <img src='assets/logo.png' alt='' />
          </div>
          <div className=''>
            <ul>
              <li>
                <Link to='/' style={{textDecoration: "none", color : "black"}}>Home</Link>
              </li>
              <li>
                <Link to='/'style={{textDecoration: "none", color : "black"}}>Scoreboard</Link>
              </li>
  <li className="account"><AccountCircleIcon/>
  <NavDropdown id="navbarScrollingDropdown">
          <NavDropdown.Item href="#action3">Nama</NavDropdown.Item>
          <NavDropdown.Item href="#action4">Log Out &nbsp;<LogoutIcon/></NavDropdown.Item>
        </NavDropdown>
</li>

            </ul>
          </div>
          <button className='navbar-items-icon' onClick={() => setSidebar(!sidebar)}>
          </button>
        </div>
      </header>
    </>
  )
}

export default Header
