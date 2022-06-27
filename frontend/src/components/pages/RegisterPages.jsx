import React from 'react';
import Register from '../login/Register';
import './RegisterPages.css';

const RegisterPages = () => {
  return (
    <>
      <div className="gradient">
        <Register />

        <div className="register-image">
          <img src="./assets/greating-image.png" alt="" height={400} width={400} />
        </div>
      </div>
    </>
  );
};

export default RegisterPages;
