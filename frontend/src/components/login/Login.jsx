import React from 'react';

const Login = () => {
  const [email, setEmail] = React.useState("");
  const [password, setPassword] = React.useState("");
  const [error, setError] = React.useState("");
  const [loading, setLoading] = React.useState(false);

  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true);
    setTimeout(() => {
      if (email === "" && password === "") {
        setError("");
        setLoading(false);
      } else {
        setError("Invalid email or password!!");
        setLoading(false);
      }
    }, 2000);
  }

return(
    <div>
      <div className="greetings">
         <h1> Hallo, </h1><h1> Welcome Back!</h1>
      </div>
      <div className="container-login">
        <div className='login-box'>
          <form onSubmit={handleSubmit}>
        <h2>Log In</h2>
        <label>Email</label>
        <input type="text" placeholder="Email" id="useremail" value={email} onChange={(e) => setEmail(e.target.value)}></input>
        <label>Password</label>
        <input type="password" placeholder="Password" id="password" value={password} onChange={(e) => setPassword(e.target.value)}></input>
          <div className="login-form-error">
             {error}
            </div>
        <button type="submit" value="Log In" disabled={loading}>{loading ? 'Wait' : 'Log In'}</button>
          <div className="signuplink">
        <a value>Don't have an account? </a><a className="Create-Account" href="#">Create here</a>
            </div>
          </form>
        </div>
      </div>
    </div>
    )
}

export default Login;
