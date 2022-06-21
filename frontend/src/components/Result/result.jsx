import React from "react"
import TestResult from "./resultTest"
import "./style.css"

const Result = () => {
  
  const handleClick = (e) => {
    console.log(e.target)
}

  return (
    <div className="result">
      <h1 className='title'>Result</h1>
      
      <TestResult />

      <div className="toScoreboard">
        <button onClick={handleClick} data-id='scoreboard'>Scoreboard</button>
      </div>
    </div>
  )
}

export default Result