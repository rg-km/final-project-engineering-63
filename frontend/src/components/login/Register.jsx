import React from 'react';
import { Link } from 'react-router-dom';

const Register = () => {
  const [name, setName] = React.useState('');
  const [email, setEmail] = React.useState('');
  const [password, setPassword] = React.useState('');
  const [error, setError] = React.useState('');
  const [loading, setLoading] = React.useState(false);

  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true);
    setTimeout(() => {
      if (name === '' && email === '' && password === '') {
        setError('');
        setLoading(false);
      } else {
        setError('Complete all fields!!');
        setLoading(false);
      }
    }, 2000);
  };

  return (
    <div>
      <div className="greetings-register">
        <h1> Hi, </h1>
        <h1> Welcome to</h1>
        <h1>EDUMAR!</h1>
      </div>
      <div className="container-register">
        <div className="register-box">
          <form onSubmit={handleSubmit}>
            <h2>Create new account</h2>
            <label>Name</label>
            <input type="text" placeholder="Name" id="username" value={name} onChange={(e) => setName(e.target.value)}></input>
            <label>Email</label>
            <input type="text" placeholder="Email" id="useremail" value={email} onChange={(e) => setEmail(e.target.value)}></input>
            <label>Password</label>
            <input type="password" placeholder="Password" id="password" value={password} onChange={(e) => setPassword(e.target.value)}></input>
            <div className="login-form-error">{error}</div>
            <button type="submit" value="Register" disabled={loading}>
              {loading ? 'Wait' : 'Register'}
            </button>
            <div className="signuplink">
              <a value>Already have an account? </a>
              <Link to="login">
                <a className="Create-Account" href="">
                  Login
                </a>
              </Link>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Register;
