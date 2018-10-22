import React, { Component } from 'react'
import { Route, Switch, NavLink } from 'react-router-dom'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faUpload, faSearch } from '@fortawesome/free-solid-svg-icons'

import Upload from './upload';

import 'bulma/css/bulma.css'

class App extends Component {
    render() {
        return(
            <div className="App">
                <nav className="navbar has-shadow is-spaced is-light" role="navigation">
                    <a className="navbar-item" href="/">
                        <h1 className="title">DummyJsonGenerator</h1>
                    </a>
                        <NavLink className="navbar-item" to="/">
                            <span><FontAwesomeIcon icon={faUpload} color={"hsl(141, 71%, 48%)"} /></span>
                            <span>Upload</span>
                        </NavLink>
                        <NavLink className="navbar-item" to="/search">
                            <span><FontAwesomeIcon icon={faSearch} color={"hsl(271, 100%, 71%)"} /></span>
                            <span>Search</span>
                        </NavLink>
                </nav>
                <section className="section">
                    <Switch className="container">
                        <Route exact path="/"  component={Upload} />
                        <Route exact path="/search" />
                    </Switch>
                </section>
            </div>
        )
    }
}

export default App;
