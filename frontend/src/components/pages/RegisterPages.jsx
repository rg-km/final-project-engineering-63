import React from 'react';
import Login from '../login/Login';
import GreatingsImage from '../../../public/assets/greating-image.png';

const RegisterPages = () => {
  return (
    <>
      <Login />
      <div className="login-image">
        <img src={GreatingsImage} alt="" height={400} width={400} />
      </div>
    </>
  );
};

export default RegisterPages;
