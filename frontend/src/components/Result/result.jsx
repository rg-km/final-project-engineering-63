import React from "react"
import TestResult from "./ResultTest"
import "./style.css"
import { Link } from "react-router-dom"

const Result = () => {
  
//   const handleClick = (e) => {
//     console.log(e.target)
// }

  return (
    <div className="result">
      <h1 className='title'>Result</h1>
      
      <TestResult />

      <div className="toScoreboard">
        <button><Link to= "board" style={{textDecoration: "none", color : "black"}}>Scoreboard</Link></button>
      </div>
    </div>
  )
}

export default Result