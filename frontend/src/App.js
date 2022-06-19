import React from "react"
import "./App.css"
import { BrowserRouter as Router, Switch, Route } from "react-router-dom"
// import Footer from "./components/home/Footer"
// import Header from "./components/home/header/Header"
// import HomePages from "./components/pages/HomePages"
// import LoginPages from "./components/pages/LoginPages"
import CoursePages from "./components/pages/CoursePages"

const App = () => { 
  return (
    <>
      {/* <Router>
        <Header />
        <Switch> */}
          {/* <Route path='/' exact component={HomePages} />
          <Route path='/login' exact component={LoginPages} /> */}
          {/* <Route path='/course' exact component={CoursePages} />
        </Switch>
        <Footer />
      </Router> */}
      <CoursePages />
    </>
  )
}

export default App
