import React from 'react';
import './App.css';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Footer from './components/home/Footer';
import Header from './components/home/header/Header';
import HomePages from './components/pages/HomePages';
import LoginPages from './components/pages/LoginPages';
import CoursePages from './components/pages/CoursePages';
import ScoreboardPages from './components/pages/scoreboardPages';
import RegisterPages from './components/pages/RegisterPages';
import ResultPages from './components/pages/ResultPages';

const App = () => {
  return (
    <>
      <Router>
        <Header />
        <Switch>
          {/* <CoursePages /> */}
          <Route path="/" exact component={HomePages} />
          <Route path="/login" exact component={LoginPages} />
          <Route path="/course" exact component={CoursePages} />
          <Route path="/board" exact component={ScoreboardPages} />
          <Route path="/register" exact component={RegisterPages} />
          <Route path="/result" exact component={ResultPages} />
        </Switch>
        <Footer />
      </Router>
      {/* <Result /> */}
      {/* <CoursePages /> */}
    </>
  );
};

export default App;
