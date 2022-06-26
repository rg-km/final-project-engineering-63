import React, { useState, useEffect, useRef } from 'react';
import { QuestionAnswer } from './data/QuestionAnswer';
import Form from 'react-bootstrap/Form';
import { NavItem } from 'react-bootstrap';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import "./Course.css"


const Course = () => {
  const [show, setShow] = useState(false);
  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);
  const arr = ['A', 'B', 'C', 'D'];
  const { id, question, options } = QuestionAnswer[0];
  const interval = useRef(null);
  const [stopwatch, setStopwatch] = useState('00:00:00');

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
        <div className="course-question">{question}</div>
      <Modal show={show} onHide={handleClose}>
      <Modal.Body className="login3"><p>Times Up!!!</p>
      <Button className="login2" variant="warning" onClick={handleClose}>Close
        </Button>
      </Modal.Body>
      </Modal>
      <div className="course-timer">{stopwatch}</div>
        <Form className='choice'>
          {QuestionAnswer[0].options.map((item, i) => (
              <Form.Check label={item.choice} name="group1" value={item.choice} type={'radio'} />
          ))}
          {/* {['radio'].map((type) => (
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
        </Form>
        
      </div>
      <button type="submit-answer">Submit</button>
      <button type="next">Next</button>
    </div>
  );
};

export default Course;
