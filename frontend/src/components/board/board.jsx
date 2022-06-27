import React, { useState } from 'react'
import Score from './score'
import "./style.css"

const Board = () => {

    const handleClick = (e) => {
        console.log(e.target)
    }
  
        return (
            <div className="board">
                <h1 className='scoreboard'>Scoreboard</h1>

                <div className="courses">
                    <button onClick={handleClick} data-id='course-vocab'>Vocab</button>
                    <button onClick={handleClick} data-id='course-grammar'>Grammar</button>
                    <button onClick={handleClick} data-id='course-tenses'>Tenses</button>
                </div>

                <Score />

                {/* <div className='back'>
                    <button onClick={handleClick} data-id='toHome'>Back to Home
                    </button>
                </div> */}
            </div>
  )
}

export default Board