import React, { useState, useEffect, useRef } from 'react';
import { QuestionAnswer } from "./data/Question&answer"

const Course = () => {

const {id, question, options} = QuestionAnswer[0]
return (
    <div>
        <div className='course-overlay'>
            <div className='course-question'>{question}</div>
            <label className='course-choice'>
                {options.map(item => (
                    <div>{item.choice}</div>
                    ))}
            </label>
    </div>
    <button type="submit">Submit</button>
    <button type="next">Next</button>
    </div>
    );
};

export default Course;