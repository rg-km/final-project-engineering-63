import React, { useState, useEffect, useRef } from 'react';
import { QuestionAnswer } from './data/QuestionAnswer';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import "./Course.css"

const Course = () => {
  const [show, setShow] = useState(false);
  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);
  const interval = useRef(null);
  const [stopwatch, setStopwatch] = useState('00:00:00');
  const [currentIndex, setCurrentIndex] = useState(0);
  const { id, question, options } = QuestionAnswer[currentIndex];

  const nextQuestion = () => {
    if(QuestionAnswer.length - 1 === currentIndex)
    return;
    setCurrentIndex(currentIndex + 1);
  };

  const previousQuestion = () => {
    if(currentIndex === 0)
    return;
    setCurrentIndex(currentIndex - 1);
  };

  function getTimeRemaining(endtime) {
    const total = Date.parse(endtime) - Date.parse(new Date());
    const hours = Math.floor((total % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
    const minutes = Math.floor((total % (1000 * 60 * 60)) / (1000 * 60));
    const seconds = Math.floor((total % (1000 * 60)) / 1000);

    return {
      total,
      hours,
      minutes,
      seconds,
    };
  }

  function startCount(endtime) {
    let { total, hours, minutes, seconds } = getTimeRemaining(endtime);
    if (total >= 0) {
      setStopwatch((hours > 9 ? hours : '0' + hours) + ':' + (minutes > 9 ? minutes : '0' + minutes) + ':' + (seconds > 9 ? seconds : '0' + seconds));
    } else {
      clearInterval(interval.current);
      setShow(true)
    } 
  }

  useEffect(() => {
    const endtime = new Date(Date.parse(new Date()) + 3600 * 1000);
    interval.current = setInterval(() => startCount(endtime), 1000);
    return () => clearInterval(interval.current);
  }, []);

  return (
    <div>
      <div className="course-overlay">
        <div className="course-question">{currentIndex + 1}. {question}</div>
        <div type="radio"className="course-choice">
        <Modal show={show} onHide={handleClose}>
      <Modal.Body className="login3"><p>Times Up!!!</p>
      <Button className="login2" variant="warning" onClick={handleClose}>Close
        </Button>
      </Modal.Body>
      </Modal>
      <div className="course-timer">{stopwatch}</div>
            {options.map((item) => (
                <div><input type="radio" name="group" class="form-check-input" id="radio"/>{item.choice}</div>
            ))}
        </div>
          {/* {['radio'].map((item, type) => (
        <div className="mb-3">
          <Form.Check
            label="1"
            name="group1"
            type={type}
          />
          <Form.Check
            label="2"
            name="group1"
            type={type}
          />
          <Form.Check
            label="3"
            name="group1"
            type={type}
          />
        </div>
      ))} */}
      </div>
      <button type="submit-answer">Submit</button>
      <button disabled= {currentIndex === 0 ? true : false} type= "button-previous" onClick={() => previousQuestion()}>Previous</button>
      <button disabled= {QuestionAnswer.length - 1 === currentIndex ? true : false} type= "button-next" onClick={() => nextQuestion()}>Next</button>
    </div>
  );
};

export default Course;