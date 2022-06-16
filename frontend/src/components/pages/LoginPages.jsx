import React from "react"
import Login from "../login/Login"
import "./LoginPages.css"


const LoginPages = () => {
    return (
      <>
      <div className="gradient">
        <Login/>
          <div className="login-image">
            <img src="./assets/greating-image.png" alt="" height={400} width={400}/>
        </div>
      </div>
      </>
    )
  }
  
  export default LoginPages