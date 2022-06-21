import React from 'react'

const TestResult = () => {
  return (
        <div id="isResult">
           {Item()}
        </div>
  )
}

function Item(){
    return (
        <div className="flex">
            <div className='text'>
                <p className='isText'>Correct</p>
                <span className='correctScore'><p>13</p></span>
            </div>
            <div className='text'>
                <p className='isText'>Wrong</p>
                <span className='wrongScore'><p>2</p></span>
            </div>
            <div className='text'>
                <p className='isText'>Duration</p>
                <span className='times'><p>13:00</p></span>
            </div>
        </div>
    )
}

export default TestResult