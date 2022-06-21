import React from 'react'

const Score = () => {
  return (
        <div id="scores">
           {Item()}
        </div>
  )
}

function Item(){
    return (
        <div className="flex">
            <div className='number'>
                <p>No</p>
            </div>
            <div className='name'>
                <p>Name</p>
            </div>
            <div className='score'>
                <p>Score</p>
            </div>
            <div className='duration'>
                <p>Duration</p>
            </div>
        </div>
    )
}

export default Score