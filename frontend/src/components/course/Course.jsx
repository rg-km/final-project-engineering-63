import React, { useState, useEffect, useRef } from 'react';
import { QuestionAnswer } from "./data/Question&answer"

const Course = () => {

const {id, question, options} = QuestionAnswer[0]
const interval = useRef(null)
const [stopwatch, setStopwatch] = useState('00:00:00')

function getTimeRemaining(endtime) {
    const total = Date.parse(endtime) - Date.parse(new Date());
    const hours = Math.floor((total % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
    const minutes = Math.floor((total % (1000 * 60 * 60)) / (1000 * 60));
    const seconds = Math.floor((total % (1000 * 60)) / 1000);   

return {
    total, hours, minutes, seconds
    };
}

function startCount(endtime) {
    let { total, hours, minutes, seconds } = getTimeRemaining(endtime);
    if(total>=0){
        setStopwatch ( (hours > 9 ? hours : '0' + hours) + ':' + (minutes > 9 ? minutes : '0' + minutes) + ':' + (seconds > 9 ? seconds : '0' + seconds))
    }else{
        clearInterval(interval.current);
    }
}

useEffect(() => {
    const endtime = new Date(Date.parse(new Date()) + 3600 * 1000);
    interval.current = setInterval(() => startCount(endtime), 1000);
    return () => clearInterval(interval.current);
}, []);


return (
    <div>
        <div className='course-overlay'>
            <div className='course-question'>{question}</div>
            <label className='course-choice'>
                {options.map(item => (
                    <div>{item.choice}</div>
                    ))}
            </label>
        <div className='course-timer'>{stopwatch}</div>
    </div>
    <button type="submit-answer">Submit</button>
    <button type="next">Next</button>
    </div>
    );
};

export default Course;