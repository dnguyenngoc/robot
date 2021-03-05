import React from "react"
import { Route, Redirect } from "react-router-dom"

function userIsLoggedIn(){
    return false
}

export default function AppRouter({path, container, auth }){
    return(
        auth && !userIsLoggedIn() ? <Route path={path} render={() => (<Redirect to="/login"/>)}/> : <Route path={path} component={container} />
    )
}