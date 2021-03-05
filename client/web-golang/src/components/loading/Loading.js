import React from 'react';
import './Loading.scss'


export default class Loading extends React.Component {
    render(){
        return (
            <div className="loading">
                <div className="loader"></div>
                <div className="bar__loader">Loading ...</div>
            </div>
        )
        
    }
}

